## 1. Ensure Ansible is installed:


sudo apt install ansible

## Make sure Ansible configured properly and have proper ssh access from master to slaves.

ansible -m ping all
ansible -m ping groupname
ansible -m ping nodename


## 2. Run the Playbook:
ansible-playbook -i inventory harden_linux.yml

