Vagrant.configure("2") do |config|
 config.vm.box = "hashicorp/bionic64"
 config.vm.define "worker"
 config.vm.network "public_network", ip: "192.168.1.100", bridge: "wlan0"
 config.vm.provision "ansible" do |ansible|
   ansible.playbook = "docker.yaml"
   ansible.groups = {
     "vm" => ["worker"],
     "vm:vars" => ["ansible_python_interpreter": "/usr/bin/python3", "ansible_user": "vagrant"]
   }
 end
#  config.vm.provision "shell", path: "get-docker.sh"
end
