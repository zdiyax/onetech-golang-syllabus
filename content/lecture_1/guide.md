###Гайд как наладить себе Workspace до пятницы:

* Установить go 1.12
(проверить в терминале командой `go version`)
* Установить git
(проверить в терминале командой `git version`)
* Забить в конфиги свой username & email из GitHub командами: 

`git config --global user.name "MY_USER_NAME"`

`git config --global user.email "MY_USER_EMAIL"`

* Скачать и установить GoLand [тут](https://www.jetbrains.com/go/download/)
* Открыть GoLand и создать новый проект (название проекта прям в его пути)
* Создать файл `hello.go` с main функцией и хэлоуворлдом
* Создать в GitHub'e новый репозиторий
* В проекте в терминале начинаем:

```
git init
git add .
git remote add origin YOUR_GITHUB_REPO.git
git commit -m "SOME MEANINGFUL MESSAGE"
git push origin master
```

