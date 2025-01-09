# GoServer: 사용자 인증 및 데이터 관리를 위한 RESTful API 서버

## 프로젝트 개요

GoServer는 사용자 인증, 데이터 관리 확장 가능하고 효율적인 RESTful API 서버입니다.   
Go 언어의 강력한 성능과 동시성 기능을 기반으로 구축되었으며, Gin 웹 프레임워크를 사용하여 빠르고 안정적인 API 라우팅을 제공합니다.   
GoServer는 마이크로서비스 아키텍처에 적합하도록 설계되었으며, MongoDB, MySQL, Redis와의 유연한 연동을 통해 다양한 데이터 스토리지 요구 사항을 충족합니다.

## 핵심 기능 및 아키텍처

GoServer는 사용자 인증, 데이터 CRUD(생성, 읽기, 업데이트, 삭제) 두 가지 주요 영역에 중점을 두고 개발되었습니다.

### 1. 사용자 인증 (Authentication)

*   **JWT (JSON Web Token) 기반 인증:** 사용자의 로그인 요청을 처리하고, 인증 성공 시 JWT 토큰을 발급합니다. 이 토큰은 이후 API 요청 시 사용자 인증에 사용됩니다.
*   **보안:** `golang.org/x/crypto` 패키지를 활용하여 사용자 비밀번호를 안전하게 해싱하고 저장합니다.
*   **Redis를 이용한 토큰 관리:** Redis에 사용자 토큰을 저장하여 빠른 인증 및 세션 관리를 지원합니다. 이를 통해 로그인 상태 유지, 토큰 만료, 강제 로그아웃 등의 기능을 효율적으로 구현할 수 있습니다.

### 2. 데이터 관리 (Data Management)

*   **다중 데이터베이스 지원:** MongoDB, MySQL, Redis와의 연동을 통해 다양한 데이터 저장 방식을 지원합니다.
    *   **MongoDB:** 문서 지향 데이터베이스로, 유연한 스키마와 빠른 개발을 지원합니다. 사용자 프로필, 로그 데이터 등 비정형 데이터 저장에 적합합니다.
    *   **MySQL:** 관계형 데이터베이스로, 정형화된 데이터와 트랜잭션 처리에 강점을 가집니다. 주문 정보, 결제 내역 등 데이터 무결성이 중요한 데이터 저장에 적합합니다.
    *   **Redis:** 인메모리 데이터 구조 저장소로, 빠른 읽기/쓰기 성능을 제공합니다. 캐싱, 세션 관리, 실시간 데이터 처리에 유용합니다.
*   **CRUD API:** 각 데이터베이스에 대한 CRUD (Create, Read, Update, Delete) API 엔드포인트를 제공하여 데이터를 쉽게 관리할 수 있습니다.

## 기술 스택

*   **언어:** [Go](https://go.dev/) (1.21 이상) - 간결하고 효율적인 언어로, 동시성 프로그래밍에 강점을 가집니다.
*   **웹 프레임워크:** [Gin](https://github.com/gin-gonic/gin) - 빠르고 가벼운 웹 프레임워크로, API 라우팅 및 미들웨어 처리에 탁월합니다.
*   **데이터베이스:**
    *   [MongoDB](https://www.mongodb.com/) ([go.mongodb.org/mongo-driver](https://go.mongodb.org/mongo-driver)) - 유연한 스키마와 확장성을 제공하는 NoSQL 데이터베이스입니다.
    *   [MySQL](https://www.mysql.com/) ([github.com/go-sql-driver/mysql](https://github.com/go-sql-driver/mysql)) - 안정적이고 널리 사용되는 관계형 데이터베이스입니다.
    *   [Redis](https://redis.io/) ([github.com/redis/go-redis/v9](https://github.com/redis/go-redis/v9)) - 고성능 인메모리 데이터 저장소로, 캐싱 및 실시간 데이터 처리에 적합합니다.
*   **로깅:** [logrus](https://github.com/sirupsen/logrus) - 구조화된 로깅을 지원하는 라이브러리입니다.
*   **환경 설정:** [godotenv](https://github.com/joho/godotenv) - `.env` 파일에서 환경 변수를 로드하는 라이브러리입니다.
*   **인증:** [jwt-go](https://github.com/golang-jwt/jwt/v4) - JWT 토큰 생성 및 검증을 위한 라이브러리입니다.
*   **구성 파일:** [go-toml](https://github.com/pelletier/go-toml/v2) - TOML 형식의 구성 파일을 파싱하는 라이브러리입니다.
*   **로깅:** [lumberjack](https://github.com/natefinch/lumberjack) - 로그 파일 로테이션을 지원하는 라이브러리입니다.

## 프로젝트 구조
```
GoServer/
├── config/ # 환경 설정 관련 파일 및 로직
├── database/ # 데이터베이스 연결 및 초기화
│ ├── mongo/
│ ├── mysql/
│ └── redis/
├── handler/ # API 핸들러 (라우팅, 요청 처리)
├── logger/ # 로깅 설정 및 유틸리티
├── model/ # 데이터 모델 (데이터베이스 스키마에 대응)
├── router/ # API 라우터 설정
├── usecase/ # 비즈니스 로직 (API 핸들러와 데이터베이스 사이의 추상화 계층)
├── main.go # 메인 애플리케이션 진입점
├── go.mod # Go 모듈 의존성 관리 파일
├── go.sum # Go 모듈 체크섬
└── .env # 환경 변수 예시 파일
```

## 시작하기

### 필수 조건

*   [Go](https://go.dev/dl/) (1.21 이상) 설치
*   [MongoDB](https://www.mongodb.com/try/download/community) 설치 및 실행
*   [MySQL](https://dev.mysql.com/downloads/installer/) 설치 및 실행
*   [Redis](https://redis.io/download) 설치 및 실행

### 빌드 및 실행

1. **저장소 복제:**

    ```bash
    git clone https://github.com/WindowsHyun/GoServer
    cd GoServer
    ```

2. **의존성 설치:**

    ```bash
    go mod tidy
    ```

3. **Swagger 초기화 (옵션):**

    ```bash
    swag init --parseDependency
    ```

4. **빌드:**

    ```bash
    go build -o GoServer main.go
    ```

5. **실행:**

    ```bash
    ./GoServer
    ```

    또는 `go run main.go` 명령어를 사용하여 직접 실행할 수도 있습니다.

### 환경 설정

애플리케이션 설정은 `default.toml` 통해 관리됩니다.

*   `SERVER_PORT`: 서버가 실행될 포트 (예: `9186`)
*   `MONGO_URI`: MongoDB 연결 문자열 (예: `mongodb://localhost:27017`)
*   `MONGO_DATABASE`: MongoDB 데이터베이스 이름 (예: `GoServer`)
*   `MYSQL_USER`: MySQL 사용자 이름 (예: `root`)
*   `MYSQL_PASSWORD`: MySQL 비밀번호
*   `MYSQL_DATABASE`: MySQL 데이터베이스 이름 (예: `GoServer`)
*   `MYSQL_HOST`: MySQL 호스트 주소 (예: `localhost`)
*   `MYSQL_PORT`: MySQL 포트 (예: `3306`)
*   `REDIS_ADDRESS`: Redis 주소 (예: `localhost:6379`)
*   `REDIS_PASSWORD`: Redis 비밀번호 (선택 사항)
*   `JWT_SECRET`: JWT 서명에 사용될 비밀 키

### 데이터베이스 설정

애플리케이션을 실행하기 전에 MongoDB, MySQL, Redis 데이터베이스가 올바르게 설정되어 있고 실행 중인지 확인하세요.   
필요한 데이터베이스 및 컬렉션/테이블을 미리 생성해야 할 수도 있습니다.

## API 문서

API 문서는 별도로 제공될 예정입니다. (예: Swagger)

## 확장성 및 추가 개발

GoServer는 마이크로서비스 아키텍처에 적합하도록 설계되어, 개별 기능의 독립적인 배포 및 확장이 가능합니다.

### 1. 새로운 API 엔드포인트 추가

*   `handler/` 디렉토리에 새로운 핸들러 파일을 추가합니다.
*   `usecase/` 디렉토리에 해당 핸들러에 대한 비즈니스 로직을 구현합니다.
*   `router/router.go` 파일에 새로운 라우트를 등록합니다.

### 2. 데이터베이스 스키마 변경

*   `model/` 디렉토리에서 데이터 모델을 수정합니다.
*   `database/` 디렉토리에서 데이터베이스 초기화 로직을 업데이트합니다.

### 3. 새로운 데이터베이스 연동

*   `database/` 디렉토리에 새로운 데이터베이스 클라이언트를 추가합니다.
*   `usecase/` 디렉토리에서 새로운 데이터베이스를 사용하는 비즈니스 로직을 구현합니다.

### 4. 인증 방식 변경

*   `handler/auth.go` 파일에서 인증 로직을 수정합니다.
*   다른 인증 방식 (예: OAuth 2.0)을 구현합니다.

## 기여하기

프로젝트에 기여하고 싶으시다면, 언제든지 풀 리퀘스트를 보내주세요.   
코드 스타일 및 컨벤션을 준수해주시면 감사하겠습니다.