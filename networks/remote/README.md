# Remote 테스트넷 설명

- name : remotenet
- AWS 기준


# Terraform 설치

<pre><code>$ wget https://releases.hashicorp.com/terraform/(버전)/terraform_(버전)_linux_amd64.zip
$ unzip terraform_(버전)_linux_amd64.zip
$ sudo cp terraform /usr/local/bin
$ sudo chmod +x /usr/local/bin
$ terraform --version</code></pre>

# Ansible 설치

<pre><code>$ sudo amazon-linux-extras install ansible2</code></pre>

# AWS Access Key 생성

https://docs.aws.amazon.com/general/latest/gr/managing-aws-access-keys.html

# SSH Key 생성

<pre><code>$ export AWS_ACCESS_KEY_ID="(access key)"
$ export AWS_SECRET_ACCESS_KEY="(secret access key)"
$ export TESTNET_NAME="remotenet"
$ export CLUSTER_NAME="remotenetvalidators"
$ export SSH_PRIVATE_FILE="$HOME/.ssh/id_rsa"
$ export SSH_PUBLIC_FILE="$HOME/.ssh/id_rsa.pub"</code></pre>