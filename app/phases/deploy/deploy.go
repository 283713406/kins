package deploy

func DeployK8s() error {

	//if [ "$(check_installed deploy_k8s)" == "1" ];then return;fi
	//kydocker cp ~/.ssh deploy:/root/;
	//cd ${src_dir};
	//sleep 5;
	//make import_deploy;
	//check_error
	//docker -H tcp://0.0.0.0:2375 exec deploy bash -c "cd /home/kubespray/deploy && ansible-playbook -i inventory/kylin/inventory.ini cluster.yml -e produce_mode=true -b -v --flush-cache -f 6"

	return nil
}

func DeployApps() error {

	//if [ "$(check_installed install_kcall_apps)" == "1" ];then return;fi
	//cd ${src_dir}
	//make import_kcall;
	//check_error
	//docker -H tcp://0.0.0.0:2375 exec $kcall_name bash -c "bash /opt/kcall/scripts/kube-init.sh";
	//check_error
	//make import_aam;
	//check_error
	//docker -H tcp://0.0.0.0:2375 exec $kcall_name bash -c "cd /opt/kcall && python3 kcall.py solution deploy install all"
	//check_error

	return nil
}

func DeployBasic() error {

	//load_solution_images;
	//deploy_kcall_container;
	//installer_make_install install_bootstrap;
	//installer_make_install install_keepalived;
	//installer_make_install install_glusterfs;
	//yum_redis_cli_repo;
	//install_redis $1;
	//systemctl restart kydocker;

	return nil
}