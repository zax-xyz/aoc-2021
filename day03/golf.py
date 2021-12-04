"""
Based on original one-liner solution by "stap.le™wis" in UNSW CSESoc Discord:
(yeet := (lambda lines, also_lines, i=0: int(lines[0], 2) * int(also_lines[0], 2) if (len(lines) < 2 and len(also_lines) < 2) else yeet((list(filter(lambda x:x[i] == ('1' if sum(1 if y[i]=='1' else -1 for y in lines) >= 0 else '0'), lines)) if len(lines) > 1 else lines), (list(filter(lambda x:x[i] == ('1' if sum(1 if y[i]=='1' else -1 for y in also_lines) < 0 else '0'), also_lines)) if len(also_lines) > 1 else also_lines),i+1)))(open('q3.txt').readlines(), open('q3.txt').readlines())
"""

# using strings
print((y:=(lambda l,a,i=0:int(l[0],2)*int(a[0],2)if(len(l+a)<3)else y((b:=lambda l,c=0:[x for x in l if x[i]==(str(c^(sum(1-int(y[i])*2for y in l)<=0)))]if len(l)>1else l)(l),b(a,1),i+1)))(l:=[*open('day3.txt')],l))
# using ints and bitwise operations (1 byte shorter)
print((y:=(lambda l,a,i=2048:l[0]*a[0]if(len(l+a)<3)else y((b:=lambda l,c=0:[x for x in l if(x&i>0)==(c^(sum(i-(y&i)*2for y in l)<=0))]if len(l)>1else l)(l),b(a,1),i>>1)))(l:=[int(x,2)for x in open('day3.txt')],l))
