package get

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//This is a simple one, we only have one API call to make:
//POST /api/public/get_images
//{
//     "api_key": "{{api_key}}",
//     "api_key_secret": "{{api_key_secret}}"
// }

type Image struct {
	RestrictToNetwork       bool              `json:"restrict_to_network"`
	Memory                  int               `json:"memory"`
	ZoneName                string            `json:"zone_name"`
	XRes                    int               `json:"x_res"`
	Description             string            `json:"description"`
	ImageId                 string            `json:"image_id"`
	PersistentProfilePath   string            `json:"persistent_profile_path"`
	FriendlyName            string            `json:"friendly_name"`
	VolumeMappings          map[string]string `json:"volume_mappings"`
	RestrictToZone          bool              `json:"restrict_to_zone"`
	DockerToken             string            `json:"docker_token"`
	PersistentProfileConfig map[string]string `json:"persistent_profile_config"`
	Cores                   float64           `json:"cores"`
	DockerRegistry          string            `json:"docker_registry"`
	Available               bool              `json:"available"`
	RunConfig               struct {
		Hostname string `json:"hostname"`
	} `json:"run_config"`
	ImageAttributes  []ImageAttribute      `json:"imageAttributes"`
	DockerUser       string                `json:"docker_user"`
	RestrictToServer bool                  `json:"restrict_to_server"`
	Enabled          bool                  `json:"enabled"`
	Name             string                `json:"name"`
	ZoneId           string                `json:"zone_id"`
	YRes             int                   `json:"y_res"`
	ServerId         string                `json:"server_id"`
	NetworkName      string                `json:"network_name"`
	ExecConfig       map[string]ExecConfig `json:"exec_config"`
	Hash             string                `json:"hash"`
	ImageSrc         string                `json:"image_src"`
}

type ImageAttribute struct {
	ImageId  string `json:"image_id"`
	AttrId   string `json:"attr_id"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Value    string `json:"value"`
}

type ExecConfig struct {
	FirstLaunch map[string]FirstLaunch `json:"first_launch"`
	Go          map[string]Go          `json:"go"`
}

type FirstLaunch struct {
	Environment map[string]string `json:"environment"`
	Cmd         string            `json:"cmd"`
}

type Go struct {
	Cmd string `json:"cmd"`
}

type GetImagesResponse struct {
	Images []Image `json:"images"`
}

//Return a list of images in either simple or verbose format
func GetImages(url string, apiKey string, apiKeySecret string, notls bool, verbose bool) {
	var err error
	url = url + "/api/public/get_images"
	var jsonStr = []byte(`{"api_key":"` + apiKey + `","api_key_secret":"` + apiKeySecret + `"}`)
	var req *http.Request
	req, err = http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	//Unmarshal the response into a GetImagesResponse struct
	var getImagesResponse GetImagesResponse
	err = json.Unmarshal(body, &getImagesResponse)
	if err != nil {
		panic(err)
	}
	if verbose {
		for _, image := range getImagesResponse.Images {
			//Print every attribute of our struct, with the format:
			//Name: value
			fmt.Printf("Friendly Name: %s\n", image.FriendlyName)
			fmt.Printf("Image Name: %s\n", image.Name)
			fmt.Printf("Image ID: %s\n", image.ImageId)
			fmt.Printf("Image Src: %s\n", image.ImageSrc)
			fmt.Printf("Image Hash: %s\n", image.Hash)
			fmt.Printf("Image Attributes:\n")
			for _, attribute := range image.ImageAttributes {
				fmt.Printf("\tName: %s\n", attribute.Name)
				fmt.Printf("\tCategory: %s\n", attribute.Category)
				fmt.Printf("\tValue: %s\n", attribute.Value)
			}
			fmt.Printf("Exec Config:\n")
			for key, value := range image.ExecConfig {
				fmt.Printf("\t%s: %s\n", key, value)
			}
			fmt.Printf("Volume Mappings:\n")
			for key, value := range image.VolumeMappings {
				fmt.Printf("\t%s: %s\n", key, value)
			}
			fmt.Printf("Persistent Profile Config:\n")
			for key, value := range image.PersistentProfileConfig {
				fmt.Printf("\t%s: %s\n", key, value)
			}
			fmt.Printf("Persistent Profile Path: %s\n", image.PersistentProfilePath)
			fmt.Printf("Restrict To Zone: %t\n", image.RestrictToZone)
			fmt.Printf("Restrict To Network: %t\n", image.RestrictToNetwork)
			fmt.Printf("Restrict To Server: %t\n", image.RestrictToServer)
			fmt.Printf("Enabled: %t\n", image.Enabled)
			fmt.Printf("Available: %t\n", image.Available)
			fmt.Printf("Cores: %f\n", image.Cores)
			fmt.Printf("Docker Registry: %s\n", image.DockerRegistry)
			fmt.Printf("Docker User: %s\n", image.DockerUser)
			fmt.Printf("Server ID: %s\n", image.ServerId)
			fmt.Printf("Zone ID: %s\n", image.ZoneId)
			fmt.Printf("Network Name: %s\n", image.NetworkName)
			fmt.Printf("Y Res: %d\n", image.YRes)
			fmt.Printf("\n")

		}
	} else {
		for _, image := range getImagesResponse.Images {
			fmt.Println("Friendly Name: " + image.FriendlyName)
			fmt.Println("Image Name: " + image.Name + "\n")
		}
	}

}
