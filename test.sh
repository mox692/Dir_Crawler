STR="I,am,a,pen"

STR2="icxzsdffsdasfsd am jon dfsa dfsa"

# split
ARR=(${STR//,/ })

ARR2=(${STR2// / })

# 配列の1番目の要素
echo ${ARR[0]}

# 配列の2番目の要素
echo ${ARR[1]}

echo ${ARR2[0]}