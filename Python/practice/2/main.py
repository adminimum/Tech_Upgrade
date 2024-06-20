import argparse
import json
import os
from datetime import datetime

def prebuild_read_log_keys():
    
    pass

def read_arguments():
    parser = argparse.ArgumentParser(prog="ReadLogAndAnalyze",
                                     description="Reading and analyzing log program")
    
    parser.add_argument("-f","--logpath",required=True,
                        help="Path for log file")
    parser.add_argument("-o","--outpath",
                        help="Path for program output file")
    parser.add_argument("-p","--outprint",action="store_true", default=False,
                        help="Print result to stand output?")
    parser.add_argument("-g","--groupby",choices=['IP','ContentRequest','Days','Codes'],
                        help="Group output by [IP,ContentRequest,Days,Codes]")
    parser.add_argument("-s","--suspicious",action="store_true",
                        help="Print suspicious elements?",default=False)
    
    args = parser.parse_args()
    return args

def print_result(args,content):
    if args.outprint or (not args.outprint and args.outpath == None):
        print(content)
    if args.outpath != None:
        with open(args.outpath,"a") as file:
            file.write(str(content))

def find_suspicious(args,agent="urlgrabber"):
    
    if not os.path.exists(args.logpath):
        print("Error with opening log file ", args.logpath)
        exit(1)
    file_log = open(args.logpath,"r").readlines()
    dict_array_log = [json.loads(log_string) for log_string in file_log]
    
    # Find request with special agent
    susp_agents = {}
    for note in dict_array_log:
        if agent in note['agent']:
            if not note['agent'] in susp_agents.keys():
                susp_agents[note['agent']] = []
                susp_agents[note['agent']].append(note['remote_ip'])
                continue
            susp_agents[note['agent']].append(note['remote_ip'])
    print_result(args, f"####SUSPICIOUS agents with {agent}####")
    print_result(args, susp_agents)
    
    # Find IPs with more than 1000 requests with 404 response
    ips_and_response = {}
    for note in dict_array_log:
        if note["remote_ip"] not in ips_and_response.keys():
            ips_and_response[note["remote_ip"]] = {}
        if note["response"] not in ips_and_response[note["remote_ip"]].keys():
            ips_and_response[note["remote_ip"]][note["response"]] = 1
        else:
            ips_and_response[note["remote_ip"]][note["response"]] += 1
    
    suspicious_ips = []
    for note in dict_array_log:
        if 404 in ips_and_response[note["remote_ip"]].keys():
            if ips_and_response[note["remote_ip"]][404] >= 1000 and \
                note["remote_ip"] not in suspicious_ips:
                suspicious_ips.append(note["remote_ip"])
    print_result(args, f"####SUSPICIOUS IPs with more than 1000 requests with 404 response####")
    print_result(args, suspicious_ips)


def read_log(args):
    
    if not os.path.exists(args.logpath):
        print("Error with opening log file ", args.logpath)
        exit(1)
    file_log = open(args.logpath,"r").readlines()
    dict_array_log = [json.loads(log_string) for log_string in file_log]
    
    # IPs and request counts
    ips_count = {}
    for note in dict_array_log:
        if not note['remote_ip'] in ips_count.keys():
            ips_count[note['remote_ip']] = 1
            continue
        ips_count[note['remote_ip']] = ips_count[note['remote_ip']] + 1
    print_result(args, "####IPs counts####")
    print_result(args, ips_count) 
    
    # Response codes and requests
    resps = {}
    for note in dict_array_log:
        if not note['response'] in resps.keys():
            resps[note['response']] = 1
            resps[str(note['response'])+" requests"] = []
            if note['request'] not in resps[str(note['response'])+" requests"]:
                resps[str(note['response'])+" requests"].append(note['request'])
            continue
        resps[note['response']] = resps[note['response']] + 1
        if note['request'] not in resps[str(note['response'])+" requests"]:
            resps[str(note['response'])+" requests"].append(note['request'])
    print_result(args, "####Response and content####")
    print_result(args, resps)
    
    # Sort by date
    dates = {}
    for note in dict_array_log:
        date_var = note["time"]
        date_var = date_var.split(":")[0]

        if not date_var in dates.keys():
            dates[date_var] = 1
            dates[date_var+" requests"] = []
            if note['request'] not in dates[date_var+" requests"]:
                dates[date_var+" requests"].append(note['request'])
            continue
        dates[date_var] = dates[date_var] + 1
        if note['request'] not in dates[date_var+" requests"]:
            dates[date_var + " requests"].append(note['request'])
    print_result(args, "####Dates requests####")
    print_result(args, dates)
    
    if args.suspicious:
        find_suspicious(args)

if __name__ == "__main__":
    args = read_arguments()
    print(args)
    read_log(args)