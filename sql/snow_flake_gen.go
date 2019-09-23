package sql

const (
	 ServiceAPIPlatForm = iota
	 ServiceUser
	 ServiceVmproject
	 ServiceVmdeploy
	 ServiceParentProject
	 ServiceK8sCluster
	 ServiceK8sProject
	 ServiceHost
	 ServiceGateway
	 ServiceFileServer
	 ServiceComponent
	 ServiceCI
)

func GenerateServiceIdentities(service, length int)(identities []int64, err error) {
	worker, err := NewNode(int64(service))
	ch := make(chan int64)
	for i := 0; i < length; i++ {
		go func() {
			id := worker.Generate()
			ch <- id
		}()
	}
	identities = make([]int64, 0)
	for i := 0; i < length; i++ {
		id := <-ch
		identities = append(identities, id)
	}
	return
}