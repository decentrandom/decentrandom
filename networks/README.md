# Networks

AWS 기준 풀노드 운영시 참조하시기 바랍니다.

## 기본 설정

- 현재 테스트넷 체인 ID는 mssp-0001 입니다.

## Terraform Owners 문제

Terraform의 describe-image 요청 시 owners를 꼭 입력하도록 바뀌었습니다. aws cli에서 미리 검색 후 입력하는 형태로 해결 했습니다.

    aws ec2 describe-images --filters "Name=name,Values=CentOS Linux 7 x86_64 HVM EBS 1703_01" --region=us-east-2

CentOS Linux 7 x86_64 HVM EBS 1704_01를 지원하지 않는 지역이 있다고 하니, 1703_01로 그냥 통일했습니다.