## GO Repo

This is my personal GO repo, the place where I use to create random stuff, from prod-ready microservices to single file study services 📓.

## 🔒 Auth API

This is a small microservice to handle authentications via API, and also, my first ever project using GO 💙. Is being hosted in my personal [K8s cluster](https://.postman.co/workspace/My-Workspace~1e261ab2-0881-450b-8097-7475a6233f02/collection/39658714-7fe50393-78c3-471c-8181-3ec6713c0238?action=share&creator=39658714&active-environment=39658714-43d74b6a-a0e8-4fdb-9d10-ee1090ff00aa) (a.k.a old laptop)

[<img src="https://run.pstmn.io/button.svg" alt="Run In Postman" style="width: 128px; height: 32px;">](https://app.getpostman.com/run-collection/39658714-7fe50393-78c3-471c-8181-3ec6713c0238?action=collection%2Ffork&source=rip_markdown&collection-url=entityId%3D39658714-7fe50393-78c3-471c-8181-3ec6713c0238%26entityType%3Dcollection%26workspaceId%3Da4a05acf-ef01-4fab-9091-553b8311139f)

## 📝 PR Description

A small PR Description creator based on diffs.

Usage:

```sh
bin/prdescription --open-ai-key="your key" --branch="branch to compare agains ex.: main"
```

How to build:

```sh
go build -o bin/prdescription ./cmd/prdescription/main.go
```
