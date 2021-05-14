package utils

// GetIP gets the IP of this service
func GetIP() (string, error){
	ip := map[string]string{}
	err := FileToInterface("service.json", &ip)
	if err != nil {
		return "", err
	}
	return ip["ip"], nil
}