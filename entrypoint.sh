#!/bin/ash
set -eo pipefail

[[ "${DEBUG}" == true ]] && set -x

initialize_system() {
  echo "Initializing lepus container ..."

  # main
  debug=${debug:-1}

  # mysql config
  mysql_host=${mysql_host:-lepus.mysql}
  mysql_port=${mysql_port:-3306}
  mysql_user=${mysql_user:-root}
  mysql_password=${mysql_password:-mypassword}
  mysql_database=${mysql_database:-lepus_db}

  # redis config
  redis_host=${redis_host:-lepus.redis}
  redis_port=${redis_port:-6379}
  redis_pass=${redis_pass:-mypassword}

  # nsq config
  nsq_server=${nsq_server:-lepus.nsq\:4150}

  # mail config
  mail_host=${mail_host:-smtp.163.com}
  mail_port=${mail_port:-465}
  mail_user=${mail_user:-alarm@163.com}
  mail_pass=${mail_pass:-password}
  mail_from=${mail_from:-alarm@163.com}

  # configure etc file
  sed -i "s/\(^\s*debug\s*=\).*/\1$debug/" /app/lepus/etc/*.ini
  sed -i "s/\(^\s*mysql_host\s*=\).*/\1$mysql_host/" /app/lepus/etc/*.ini
  sed -i "s/\(^\s*mysql_port\s*=\).*/\1$mysql_port/" /app/lepus/etc/*.ini
  sed -i "s/\(^\s*mysql_user\s*=\).*/\1$mysql_user/" /app/lepus/etc/*.ini
  sed -i "s/\(^\s*mysql_password\s*=\).*/\1$mysql_password/" /app/lepus/etc/*.ini
  sed -i "s/\(^\s*mysql_database\s*=\).*/\1$mysql_database/" /app/lepus/etc/*.ini
  sed -i "s/\(^\s*redis_host\s*=\).*/\1$redis_host/" /app/lepus/etc/*.ini
  sed -i "s/\(^\s*redis_port\s*=\).*/\1$redis_port/" /app/lepus/etc/*.ini
  sed -i "s/\(^\s*redis_pass\s*=\).*/\1$redis_pass/" /app/lepus/etc/*.ini
  sed -i "s/\(^\s*nsq_server\s*=\).*/\1$nsq_server/" /app/lepus/etc/*.ini
  sed -i "s/\(^\s*mail_host\s*=\).*/\1$mail_host/" /app/lepus/etc/*.ini
  sed -i "s/\(^\s*mail_port\s*=\).*/\1$mail_port/" /app/lepus/etc/*.ini
  sed -i "s/\(^\s*mail_user\s*=\).*/\1$mail_user/" /app/lepus/etc/*.ini
  sed -i "s/\(^\s*mail_pass\s*=\).*/\1$mail_pass/" /app/lepus/etc/*.ini
  sed -i "s/\(^\s*mail_from\s*=\).*/\1$mail_from/" /app/lepus/etc/*.ini
}

check_database() {
  echo "Attempting to connect to database ..."
  prog="mysqladmin -h ${mysql_host} -P ${mysql_port} -u ${mysql_user} ${mysql_password:+-p$mysql_password} status"
  timeout=60
  while ! ${prog} >/dev/null 2>&1; do
    timeout=$((timeout - 1))
    if [[ "$timeout" -eq 0 ]]; then
      echo
      echo "Could not connect to database server! Aborting..."
      exit 1
    fi
    echo -n "."
    sleep 1
  done
  echo
}

checkdbinitmysql() {
  table=users
  if [[ "$(mysql -N -s -h "${mysql_host}" -P "${mysql_port}" -u "${mysql_user}" "${mysql_password:+-p$mysql_password}" "${mysql_database}" -e \
    "select count(*) from information_schema.tables where \
            table_schema='${mysql_database}' and table_name='${table}';")" -eq 1 ]]; then
    echo "Table ${table} exists! ..."
  else
    echo "Table ${table} does not exist! init db..."
    init_db
  fi

}

init_db() {
  echo "Initializing lepus database ..."
  #  mysql -s -h "${mysql_host}" -P "${mysql_port}" -u "${mysql_user}" "${mysql_password:+-p$mysql_password}" -e "CREATE DATABASE lepus_db;"
  mysql -s -h "${mysql_host}" -P "${mysql_port}" -u "${mysql_user}" "${mysql_password:+-p$mysql_password}" "${mysql_database}" </app/lepus/init_table.sql
  mysql -s -h "${mysql_host}" -P "${mysql_port}" -u "${mysql_user}" "${mysql_password:+-p$mysql_password}" "${mysql_database}" </app/lepus/init_data.sql
  checkdbinitmysql
}

start_system() {
  initialize_system
  check_database
  checkdbinitmysql
  echo "Starting lepus! ..."
  /usr/bin/supervisord -n -c /etc/supervisor/supervisord.conf
}

start_system

exit 0
