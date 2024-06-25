package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

type Config struct {
	Github struct {
		Token          string `json:"token"`
		ApiUrl         string `json:"api_url"`
		DefNameSpaceId int    `json:"default_namespace_id"`
	} `json:"github"`
	Gitlab struct {
		Token  string `json:"token"`
		ApiUrl string `json:"api_url"`
	} `json:"gitlab"`
}

type Project struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	Description     string `json:"description"`
	WebURL          string `json:"web_url"`
	DefaultBranch   string `json:"default_branch"`
	SSHURLToRepo    string `json:"ssh_url_to_repo"`
	HTTPURLToRepo   string `json:"http_url_to_repo"`
	VisibilityLevel string `json:"visibility"`
}

func find_project_name_by_id(ida int, config Config) string {
	auth := config.Gitlab.Token
	URL := config.Gitlab.ApiUrl + "/projects/" + strconv.Itoa(ida)
	req, _ := http.NewRequest("GET", URL, nil)
	req.Header.Add("PRIVATE-TOKEN", auth)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error on response.\n[ERROR] -", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
	}

	if resp.StatusCode != 200 {
		fmt.Println("Something happened, maybe ID is not correct, or you don't have access")
		os.Exit(1)
	}

	resp_repository := []byte(body)
	var json_projs Project
	json.Unmarshal(resp_repository, &json_projs)

	res_proj := "|" + json_projs.SSHURLToRepo + "|"

	return res_proj
}

func parsing_config() Config {
	data, err := os.ReadFile(conf_file_path)
	if err != nil {
		fmt.Println("[-] Error opening config file")
		fmt.Println(err)
		os.Exit(1)
	}
	var config_json Config
	json_err := json.Unmarshal(data, &config_json)
	if err != nil {
		fmt.Println("[-] Error parsing config file")
		fmt.Println(json_err)
		os.Exit(1)
	}
	return config_json
}

func gitlab_list(config Config) []string {
	auth := config.Gitlab.Token
	URL := config.Gitlab.ApiUrl + "/projects?min_access_level=50"
	req, _ := http.NewRequest("GET", URL, nil)
	req.Header.Add("PRIVATE-TOKEN", auth)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error on response.\n[ERROR] -", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
	}
	resp_repositories := []byte(body)
	var json_projs []Project
	json.Unmarshal(resp_repositories, &json_projs)
	var return_list []string
	for _, proj := range json_projs {
		print_proj := strconv.Itoa(proj.ID) + "\t\t| " + proj.SSHURLToRepo + " | " + proj.Name
		return_list = append(return_list, print_proj)

	}
	return return_list
}

func gitlab_create(config Config) {
	if repositories == "" || description == "" || path == "" || namespace_id == 0 {
		fmt.Println("[-] Parameters for creating repository required [description,path,namespace_id,repository] ")
		os.Exit(1)
	}

	URL_Post := config.Gitlab.ApiUrl + "/projects"
	post_data_string := fmt.Sprintf(`{
	"name": "%s",
	"description": "%s",
	"path": "%s",
	"namespace_id": "%d",
	"initialize_with_readme": "true"
	}`, repositories, description, path, namespace_id)

	fmt.Println("Body of request:")
	fmt.Println(string(post_data_string))
	var post_data = []byte(post_data_string)
	request, _ := http.NewRequest("POST", URL_Post, bytes.NewBuffer(post_data))
	request.Header.Add("PRIVATE-TOKEN", config.Gitlab.Token)
	request.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if res.StatusCode != http.StatusCreated {
		fmt.Printf("Failed to create repository. Status: %d\n", res.StatusCode)
		fmt.Println("Response body:", string(body))
		os.Exit(1)
	}

	fmt.Println("Repository has been successfully created!")
}

func gitlab_delete(config Config) {
	if project_id == 0 {
		fmt.Println("Please enter project id for deleting [-i]")
		os.Exit(1)
	}

	proj_name := find_project_name_by_id(project_id, config)

	var input string
	for {
		fmt.Printf("Are you really want to delete %s [y/n]?\n", proj_name)
		fmt.Scanln(&input)
		if input == "n" {
			fmt.Println("Good luck")
			os.Exit(0)
		} else if input == "y" {
			break
		}
	}

	URL_delete := config.Gitlab.ApiUrl + "/projects/" + strconv.Itoa(project_id)

	req, _ := http.NewRequest("DELETE", URL_delete, nil)
	req.Header.Add("PRIVATE-TOKEN", config.Gitlab.Token)
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error while executing request to delete project")
		fmt.Println(err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if resp.StatusCode != 202 {
		fmt.Println("Failure to delete project. Permission or another error")
		fmt.Println("Response body:", string(body))
		os.Exit(1)
	}
	fmt.Println("Successfully deleted")
}

func gitlab_copy(config Config) {
	// Url + projects + id + fork
	// body = path + name + namespace_id

	if repositories == "" || path == "" || project_id == 0 || namespace_id == 0 {
		fmt.Println("[-] Parameters for creating repository required [path,namespace_id,repository(name),project_id] ")
		os.Exit(1)
	}

	repository_name := find_project_name_by_id(project_id, config)

	if repository_name == "" {
		fmt.Println("Error, incorrect ID of project")
		os.Exit(1)
	}

	URL := config.Gitlab.ApiUrl + "/projects/" + strconv.Itoa(project_id) + "/fork"
	body := fmt.Sprintf(`{
	"name": "%s",
	"path": "%s",
	"namespace_id": "%d"}`, repositories, path, namespace_id)
	body_req := []byte(body)

	fmt.Println("Body of request:")
	fmt.Println(string(body))

	req, _ := http.NewRequest("POST", URL, bytes.NewBuffer(body_req))
	req.Header.Add("PRIVATE-TOKEN", config.Gitlab.Token)
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error with operating reqest")
		fmt.Println(err)
		fmt.Println("Code response", resp.StatusCode)
	}

	body_resp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if resp.StatusCode != http.StatusCreated {
		fmt.Printf("Failed to create repository. Status: %d\n", resp.StatusCode)
		fmt.Println("Response body:", string(body_resp))
		os.Exit(1)
	}

	fmt.Println("Repository successfully copied")

	defer resp.Body.Close()
}

func gitlab_rename(config Config) {

	if repositories == "" || project_id == 0 {
		fmt.Println("[-] Parameters for creating repository required [repository(new_name),project_id] ")
		os.Exit(1)
	}

	URL := config.Gitlab.ApiUrl + "/projects/" + strconv.Itoa(project_id)

	body := fmt.Sprintf(`{
	"name": "%s"}`, repositories)
	body_byte := []byte(body)

	req, _ := http.NewRequest("PUT", URL, bytes.NewBuffer(body_byte))
	req.Header.Add("PRIVATE-TOKEN", config.Gitlab.Token)
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error with operating reqest")
		fmt.Println(err)
		fmt.Println("Code response", resp.StatusCode)
	}

	body_resp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if resp.StatusCode != 200 {
		fmt.Printf("Failed to rename repository. Status: %d\n", resp.StatusCode)
		fmt.Println("Response body:", string(body_resp))
		os.Exit(1)
	}

	fmt.Println("Repository renamed successfully")

	defer resp.Body.Close()

}

func gitlab_procces() {

	var config = parsing_config()
	switch action {
	case "list":
		list_projs := gitlab_list(config)
		for _, proj := range list_projs {
			fmt.Println(proj)
		}
	case "create":
		gitlab_create(config)
	case "delete":
		gitlab_delete(config)
	case "copy":
		gitlab_copy(config)
	case "rename":
		gitlab_rename(config)
	}
}

var gitlabCmd = &cobra.Command{
	Use:   "gitlab",
	Short: "Usage gitlab command",
	Run: func(cmd *cobra.Command, args []string) {
		if !validActions[action] {
			cmd.Help()
			fmt.Println("Incorrect value for ACTION")
			os.Exit(1)
		}
		gitlab_procces()
	},
}
