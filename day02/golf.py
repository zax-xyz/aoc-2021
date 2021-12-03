# p1
h=d=0
for x,y in map(str.split,open('day2.txt')):z=int(y);d+=z*(a:=len(x)-3)*(a<4);h+=z*(a>1)
print(h*d)

# p2
a=h=d=0
for x,y in map(str.split,open('day2.txt')):
 z=int(y)
 if'f'!=x[0]:a+=z*(len(x)-3)
 else:h+=z;d+=z*a
print(h*d)
