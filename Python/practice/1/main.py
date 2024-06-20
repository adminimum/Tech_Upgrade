import psycopg2.sql
import boto3
import psycopg2
import xml.etree.ElementTree as ET
from datetime import date
import os

config_path = "./config.xml"

def find_attrs(xml_el):
    dict_xml_el = {}
    for item in xml_el:
        if item.findall("./") != []:
            dict_xml_el[f"{item.tag}"] = find_attrs(item)
        else:            
            dict_xml_el[f"{item.tag}"] = item.text
    return dict_xml_el
        
def xml_data_parser():
    tree = ET.parse(config_path)
    root = tree.getroot()
    
    data_params = find_attrs(root)
    return data_params

def check_connection_s3(config):
    session = boto3.Session(
        aws_access_key_id=config["s3user"]["accesskey"],
        aws_secret_access_key=config["s3user"]["secretkey"])
    
    ss3 = session.resource("s3")
    buck = ss3.Bucket(config["s3user"]["s3name"])
    
    try:
        buck.objects.all()
        return True
    except:
        return False

def send_backup_s3(config,filename):
    
    ff_path = os.path.abspath(filename)
    
    session = boto3.Session(
        aws_access_key_id=config["s3user"]["accesskey"],
        aws_secret_access_key=config["s3user"]["secretkey"])
    
    ss3 = session.resource("s3")
    buck = ss3.Bucket(config["s3user"]["s3name"])
    buck.upload_file(ff_path,filename)
    
    os.remove(filename)
  
def create_db_backup_and_send(config):
    
    temp_file_name = f"backup_{date.today().day}_{date.today().month}_{date.today().year}.csv"
    
    db_config = {}
    db_config["host"] = config["dbdata"]["host"]
    db_config["database"] = config["dbdata"]["dbname"]
    db_config["user"] = config["usersql"]["username"]
    db_config["password"] = config["usersql"]["userpass"]
    
    try:
        with psycopg2.connect(**db_config) as conn:
            cursor = conn.cursor()
            print("Connected to the DB!")
            cursor.execute("""SELECT table_name, 'table' as entity_type
                              FROM information_schema.tables 
                              WHERE table_schema='public' AND table_type='BASE TABLE'
                              UNION
                              SELECT table_name, 'view' as entity_type
                              FROM information_schema.tables 
                              WHERE table_schema='public' AND table_type='VIEW'
                              UNION
                              SELECT sequence_name as table_name, 'sequence' as entity_type
                              FROM information_schema.sequences 
                              WHERE sequence_schema='public';""")
            tables = cursor.fetchall()
            
            with open(temp_file_name,"w") as f:
                for table in tables:
                    table_name = table[0]
                    backup_note = psycopg2.sql.SQL("COPY (SELECT * FROM {}) TO STDOUT WITH CSV HEADER").format(psycopg2.sql.Identifier(table_name))
                    cursor.copy_expert(backup_note,f)
            cursor.close()
    except Exception as err:
        print(err)        
    
    send_backup_s3(config,temp_file_name)

if __name__ == "__main__":
    
    xml_config_data = xml_data_parser()
    
    if not check_connection_s3(xml_config_data):
        print("Error to connect s3 bucket!!!")
        exit(1)
    
    create_db_backup_and_send(xml_config_data)
    print("Succesful!")