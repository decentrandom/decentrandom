# Remote 테스트넷 설명

- name : remotenet
- AWS 기준

## Terraform 설치

    wget https://releases.hashicorp.com/terraform/(버전)/terraform_(버전)_linux_amd64.zip
    unzip terraform_(버전)_linux_amd64.zip
    sudo cp terraform /usr/local/bin
    sudo chmod +x /usr/local/bin
    terraform --version

## Ansible 설치

    sudo amazon-linux-extras install ansible2

## AWS Access Key 생성

[https://docs.aws.amazon.com/general/latest/gr/managing-aws-access-keys.html](https://docs.aws.amazon.com/general/latest/gr/managing-aws-access-keys.html)

## EXPORT 추가

    export AWS_ACCESS_KEY_ID="(access key)"
    export AWS_SECRET_ACCESS_KEY="(secret access key)"
    export TESTNET_NAME="remotenet"
    export CLUSTER_NAME="remotenetvalidators"
    export SSH_PRIVATE_FILE="$HOME/.ssh/id_rsa"
    export SSH_PUBLIC_FILE="$HOME/.ssh/id_rsa.pub"

## Build

    go get -u -v github.com/decentrandom/decentrandom
    make build-linux</code></pre>

## Remote Network 생성

    SERVERS=1 REGION_LIMIT=1 make validators-start