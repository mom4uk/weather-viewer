### Statuses
[![Maintainability](https://api.codeclimate.com/v1/badges/887ca1f56e5d2ad12124/maintainability)](https://codeclimate.com/github/mom4uk/php-project-48/maintainability)
[![Github Action Status](https://github.com/mom4uk/php-project-48/actions/workflows/github-actions.yaml/badge.svg)](https://github.com/mom4uk/php-project-48/actions)
[![Test Coverage](https://api.codeclimate.com/v1/badges/887ca1f56e5d2ad12124/test_coverage)](https://codeclimate.com/github/mom4uk/php-project-48/test_coverage)

## Description
Веб-приложение для просмотра текущей погоды. Пользователь может зарегистрироваться и добавить в коллекцию одну или несколько локаций (городов, сёл, других пунктов), после чего главная страница приложения начинает отображать список локаций с их текущей погодой.

## Requirements
- Docker
- Go >= 1.26.1

## Installation

```
git clone git@github.com:mom4uk/weather-viewer.git

cd weather-viewer

make dev

make migrate
```

## Tests
Для запуска тестов выполнить:

```
make test-local
```