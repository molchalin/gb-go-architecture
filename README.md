1. Клонируем репозиторий
```
git clone git@github.com:YOUR_GITHUB/gb-go-architecture.git
```

1. Добавляем репозиторий преподавателя, чтобы в случае изменений их можно было притянуть к себе:

```
git remote add upstream https://github.com/Barugoo/gb-go-architecture.git
```

1. Создаем новую ветку
```
git checkout -b 'Homework-1'
```

1. Выбираем измения с помощью git add
по одному файлу
```
git add lesson1/shop.main.go
```
или все измененные файлы
```
git add .
```

1. Делаем коммит:
```
git commit -m 'Homework-1'
```


1. Пушим коммиты на сервер:
```
git push
```

1. Если в репозитории преподавателя возникли изменения, и вы хотите притянуть их к себе:
```
git fetch upstream
git merge upstream/master master
```