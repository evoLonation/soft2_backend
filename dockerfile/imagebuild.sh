module=$(pwd | ls | awk '$NF!="outex" && $NF!="regression.sh" {print $NF}')

for i in ${module[*]}
do
   image=${i:11}
   echo "start docker build $image"
   docker build ../ -f $i -t $image &
done