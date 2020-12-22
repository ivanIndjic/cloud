Vagrant.configure("2") do |config|
 config.vm.box = "hashicorp/bionic64"
 config.vm.define "master"
 config.vm.network "public_network", ip: "192.168.1.100", bridge: "wlan0"
 config.vm.provision "ansible" do |ansible|
   ansible.playbook = "docker.yaml"
   ansible.groups = {
     "vm" => ["master"],
     "vm:vars" => ["ansible_python_interpreter": "/usr/bin/python3", "ansible_user": "vagrant"]
   }
 end
#  config.vm.provision "shell", path: "get-docker.sh"
