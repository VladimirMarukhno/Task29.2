Задание 2. Graceful shutdown
Цель задания
Научиться правильно останавливать приложения.

Что нужно сделать
В работе часто возникает потребность правильно останавливать приложения. Например, когда наш сервер обслуживает соединения, а нам хочется, чтобы все текущие соединения были обработаны и лишь потом произошло выключение сервиса. Для этого существует паттерн graceful shutdown.

Напишите приложение, которое выводит квадраты натуральных чисел на экран, а после получения сигнала ^С обрабатывает этот сигнал, пишет «выхожу из программы» и выходит.

Советы и рекомендации
Для реализации данного паттерна воспользуйтесь каналами и оператором select с default-кейсом.

Что оценивается
Код выводит квадраты натуральных чисел на экран, после получения ^С происходит обработка сигнала и выход из программы.

Как отправить задание на проверку
Выполните задание в файле вашей среды разработки и пришлите ссылку на архив с вашим проектом через форму ниже.