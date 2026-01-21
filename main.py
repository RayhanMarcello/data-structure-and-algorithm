arr = [4,23,4,51,1]
# minValArr = arr[4]

# for i in arr :
#     if i < minValArr:
#         minValArr == i

# arr.append(minValArr)
# print("lowest", minValArr)
# print("ga di sorting", arr)

# sort

# arr.sort()
# print(arr)

# function sort
# def sort(arr_sort):
    # for i in range(len(arr_sort))






# sort manual
total = 0

# for i in range (3,6):
#     print(i, "\n")

# for j in range (0, 11, 2):
#     print(j, "\n")

# for k in range (5, -1, -1):
#     print(k, "\n")

arr_sort = [5,3,8,6,7,2]
for i in range(len(arr_sort)-1,0,-1):
    print("perulangan i: ", i)
    for j in range(i):

        print("perlangan j" , j+1)

        # print("perulangan j: ", j+1)
        # if arr_sort[j] > arr_sort[j+1]:
        #     temp = arr_sort[j]
        #     arr_sort[j] = arr_sort[j+1]
        #     arr_sort[j+1] = temp

# print(arr_sort)


# test sendiri

for x in range(len(arr)-1, 0, -1):
    for y in range(x):
        if arr[y] > arr[y+1]:
            temp = arr[y]
            arr[y] = arr[y+1]
            arr[y+1] = temp
print(arr)

