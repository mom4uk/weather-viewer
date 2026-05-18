### Statuses
[![Maintainability](https://qlty.sh/gh/mom4uk/projects/weather-viewer/maintainability.svg)](https://qlty.sh/gh/mom4uk/projects/weather-viewer)
[![Code Coverage](https://qlty.sh/gh/mom4uk/projects/weather-viewer/coverage.svg)](https://qlty.sh/gh/mom4uk/projects/weather-viewer)

## Description
Веб-приложение для просмотра текущей погоды. Пользователь может зарегистрироваться и добавить в коллекцию одну или несколько локаций (городов, сёл, других пунктов), после чего главная страница приложения начинает отображать список локаций с их текущей погодой.

## Requirements
- Docker
- Go >= 1.26.1

## Installation

```
git clone git@github.com:mom4uk/weather-viewer.git

cd weather-viewer

заполнить .env файл

make dev
```

## Tests
Для запуска тестов выполнить:

```
make compose-test
```