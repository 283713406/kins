package images

func LoadK8sImages() error {

	//if [ "$(check_installed load_k8s_images)" == "1" ];then return;fi
	//cd /mnt/k8s/;
	//bash registryserver;
	//check_error
	//mkdir -p /opt/images;
	//bash registryserver;
	//bash registryserver make-repo /root/prepare;

	return nil
}

func LoadSolutionImages() error {

	//if [ "$(check_installed load_solution_images)" == "1" ];then return;fi
	//echo "load solution images"
	//iso=`ls | grep container-solution`
	//version=`ls | grep container-solution | awk -F '-' '{printf("%s", $4)}'`;
	//mkdir -p /mnt/solution$version;
	//mkdir -p /docker/registry/solution;
	//mount $iso /mnt/solution$version;
	//cd /mnt/solution$version;
	//bash expand-mirror /docker/registry/solution;

	return nil
}