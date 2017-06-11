
cat ./test|while read sql; do
  if echo "$sql"|grep ', ' > /dev/null;then
    id=`echo "$sql"|sed 's/\(.*\)WHERE \(.*\)/\2/g'`
    echo "$sql"| sed "s/, / WHERE $id UPDATE \`prof_job\` SET /g"
  else
    echo "$sql"
  fi
done;



#|xargs sh -c "sql=\"$0\" && id=`echo \"$0\"|sed 's/\(.*\)WHERE \(.*\)/\2/g'` && echo \"$0\"| sed 's/, / WHERE $id; UPDATE \`prof_job\` SET /g'"
