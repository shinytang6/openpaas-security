package service

import (
	"fmt"

	core_v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/typed/core/v1"
)

func CreateService(serviceClient v1.ServiceInterface, service core_v1.Service)  {
	if _, err := GetService(serviceClient, service.Name); err != nil {
		if !errors.IsNotFound(err) {
			panic(err.Error())
			return
		}
		fmt.Printf("Start creating service...\n")
		result, err := serviceClient.Create(&service)
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("Service %s created!\n", result.GetObjectMeta().GetName())
	} else {
		UpdateService(serviceClient, service)
	}
}

func GetService(serviceClient v1.ServiceInterface, servicementName string) (*core_v1.Service, error) {
	service, err := serviceClient.Get(servicementName, meta_v1.GetOptions{})
	return service, err
}

func UpdateService(serviceClient v1.ServiceInterface, service core_v1.Service) (*core_v1.Service, error) {
	fmt.Printf("Start updating service...\n")
	new_service, err := serviceClient.Update(&service)
	fmt.Printf("Service %s updated!\n", service.Name)
	return new_service, err
}

func DeleteService(serviceClient v1.ServiceInterface, serviceName string) error {
	fmt.Printf("Start deleting service %s...\n", serviceName)
	err := serviceClient.Delete(serviceName, &meta_v1.DeleteOptions{})
	return err
}