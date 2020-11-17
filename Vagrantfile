Vagrant.configure("2") do |config|
 config.vm.box = "hashicorp/bionic64"
# config.vm.define "worker"
# config.vm.network "public_network", ip: "192.168.33.10"
# config.vm.provision "ansible" do |ansible|
#   ansible.playbook = "docker.yaml"
#   ansible.groups = {
#     "vm" => ["worker"]
#   }
# end
  config.vm.provision "shell", path: "get-docker.sh"
end
