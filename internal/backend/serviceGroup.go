
// DeleteService : delete a service version
func (s *ServiceGroup) DeleteServiceGroup(serviceGroup string) {
	client := newApiClient()

	response := client.action(path.Join(serviceGroupEntity, serviceGroup)+"/", "DELETE", nil)
	response.Process(true)
}
