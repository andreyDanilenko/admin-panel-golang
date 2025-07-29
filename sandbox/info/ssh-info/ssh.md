Вот полный Markdown-файл, который вы можете скопировать или скачать целиком:

```markdown
# 🚀 Полная шпаргалка по работе с удалённым сервером (SSH/Files/Archives)

## 🔑 **1. Подключение к серверу**
```bash
# Основное подключение:
ssh -i ~/.ssh/your_key admin@server_ip

# Подключение с указанием порта:
ssh -i ~/.ssh/key.pem -p 2222 admin@host

# Копирование ключа на сервер:
ssh-copy-id -i ~/.ssh/key.pub admin@ip
```

## 📂 **2. Работа с файлами**

### 🔍 Поиск файлов
```bash
# Рекурсивный поиск по имени:
find / -type f -name "*.conf" 2>/dev/null

# Поиск в домашней директории:
find ~/ -name "important_*" -mtime -7
```

### 📝 Чтение и редактирование
```bash
# Просмотр содержимого:
cat file.txt | less

# Поиск текста в файлах:
grep -rn "password" /etc/

# Редактирование:
nano /path/to/file
vim /path/to/file
```

### ✂️ Замена содержимого
```bash
# Замена текста во всех файлах:
find . -type f -exec sed -i 's/old/new/g' {} +

# Удаление строк с "DEBUG":
sed -i '/DEBUG/d' logfile.log
```

### 🗑️ Управление файлами
```bash
# Копирование с сохранением прав:
cp -rp dir/ backup/

# Пакетное переименование:
for file in *.log; do mv "$file" "${file%.log}.bak"; done

# Безопасное удаление:
shred -zu private.key
```

## 📦 **3. Работа с архивами**

### 📥 Распаковка
```bash
# Стандартные архивы:
tar -xzvf archive.tar.gz       # .tar.gz
unzip archive.zip -d target/   # .zip

# Рекурсивная распаковка:
while [ -f archive-*.tar.gz ]; do
  tar -xzvf archive-*.tar.gz && rm -f archive-*.tar.gz
done
```

### 📤 Создание архивов
```bash
# Архивирование с максимальным сжатием:
tar -czvf backup_$(date +%F).tar.gz /important_data/

# Архивирование без путей:
zip -r -j bundle.zip *.conf
```

## ⚙ **4. Управление процессами**

### 🔄 Основные команды
```bash
# Поиск процессов:
pgrep -fl "nginx"

# Дерево процессов:
pstree -ap

# Запуск в фоне:
nohup ./start.sh > output.log 2>&1 &
```

### 🛑 Остановка процессов
```bash
# Грациозная остановка:
kill -15 PID

# Принудительное завершение:
pkill -9 -f "pattern"
```

## 🔐 **5. Безопасность и права**

### 🔒 Права доступа
```bash
# Быстрое изменение прав:
chmod -R 750 ~/private/

# Изменение владельца:
chown -R admin:admin /project/
```

### 🕵️ Проверка безопасности
```bash
# Поиск файлов с SUID:
find / -perm -4000 2>/dev/null

# Проверка открытых портов:
ss -tulnp
```

## 📊 **6. Мониторинг ресурсов**

### 📈 Системная информация
```bash
# Дисковое пространство:
df -hT /home

# Использование памяти:
free -h

# Загрузка CPU:
mpstat -P ALL 1 5
```

### 🧹 Очистка места
```bash
# Поиск больших файлов:
find / -size +100M -exec ls -lh {} +

# Очистка логов:
journalctl --vacuum-time=7d
```

## 🛠️ **7. Полезные инструменты**

### 🔗 Туннелирование
```bash
# SSH-туннель:
ssh -L 8080:localhost:80 user@remote

# SCP копирование:
scp -r ~/project admin@remote:/backups/
```

### 🧩 Работа с текстом
```bash
# Подсчёт строк кода:
cloc /project/src/

# Поиск дубликатов:
fdupes -r /home/
```

## 🆘 **8. Экстренные случаи**

### 💾 Резервное копирование
```bash
# Быстрый бэкап:
rsync -avz --delete /data/ backup:/mirror/
```

### 🔄 Восстановление
```bash
# Из резервной копии:
tar -xzvf backup.tar.gz -C /restore_path/
```

> **Pro tips:**  
> - Используйте `tmux` или `screen` для долгих операций  
> - Для сложных задач создавайте скрипты в `/usr/local/bin/`  
> - Регулярно проверяйте `crontab -l` на неожиданные задачи  

```

### Как использовать:
1. Скопируйте весь блок (от ```markdown до ```)
2. Вставьте в новый файл `server_guide.md`
3. Для удобства можно сконвертировать в PDF:
   ```bash
   pandoc server_guide.md -o server_guide.pdf --template=elegant
   ```

Файл содержит:
- Все команды в одном месте
- Логическую группировку по темам
- Готовые рецепты для частых задач
- Pro-советы в конце

Можете адаптировать под свои нужды!