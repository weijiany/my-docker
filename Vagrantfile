Vagrant.configure("2") do |config|
  config.vm.define "vm" do |m|
    m.vm.box = "ubuntu/focal64"
    m.vm.provider "virtualbox" do |v|
      v.memory = 1024
      v.cpus = 1
    end

    m.vm.network :private_network, ip: "192.168.56.10"
    m.vm.provision "shell", "path": "init.sh"
    m.vm.synced_folder ".", "/home/vagrant/app"
  end
end
