# coding: Shift_JIS
import datetime

d1 = datetime.datetime.today()
print('now:', d1)

# 10���񃋁[�v�Ŗ�150�b
for n in range(1000000000):
	if n == 5000:
		print("5000")

d2 = datetime.datetime.today()
print('now:', d2)
