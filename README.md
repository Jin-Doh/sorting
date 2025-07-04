# sorting

Go로 작성된 간단한 파일 정렬 및 중복 제거 CLI 도구입니다.

## 개요

- 지정한 텍스트 파일에서 빈 줄을 제거하고,
- 중복된 줄을 제거한 뒤,
- 오름차순으로 정렬하여
- 원본 파일에 다시 저장합니다.

## 파일 구조

- `main.go` : 정렬 및 중복 제거 기능을 수행하는 Go 프로그램
- `build_and_register.sh` : Go 프로그램을 빌드하고 macOS CLI 환경에 등록하는 스크립트

## 설치 및 사용법

### 1. 빌드 및 등록

```sh
bash build_and_register.sh
```

- 위 명령을 실행하면 `sorting` 바이너리가 `/usr/local/bin`에 등록되어 터미널에서 전역으로 사용할 수 있습니다.

### 2. 사용법

```sh
sorting --file <path_to_package_list>
```

- `<path_to_package_list>` : 정렬 및 중복 제거를 수행할 텍스트 파일 경로
- 예시:

```sh
sorting --file packages.txt
```

## 동작 방식

1. 입력 파일을 읽어 빈 줄 및 공백만 있는 줄을 제거합니다.
2. 중복된 줄을 제거합니다.
3. 남은 줄을 오름차순으로 정렬합니다.
4. 결과를 원본 파일에 덮어씁니다.

## 요구 사항
- Go (1.16 이상 권장)
- macOS 환경 (다른 OS에서도 빌드 및 사용 가능)

## 참고
- `build_and_register.sh` 실행 시 sudo 권한이 필요할 수 있습니다.
- 기존 파일이 덮어써지므로, 중요한 파일은 백업 후 사용하세요.
