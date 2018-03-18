PID=`ps -eaf | grep lala-profile-server | grep -v grep | awk '{print $2}'`
if [[ "" !=  "$PID" ]]; then
  kill -9 $PID
fi
./lala-profile-server/lala-profile-server &