
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
arr = [2,23,4,51,1]

# bubble sort
for x in range(len(arr)-1, 0, -1):
    for y in range(x):
        if arr[y] > arr[y+1]:
            temp = arr[y]
            arr[y] = arr[y+1]
            arr[y+1] = temp
print(arr)

# selection sort

arr2 = [2,6,5,1,3,4]
for i in range(0, len(arr2)-1, 1):
    print(i)
    cur_min_idx = i
    for j in range(i+1, len(arr2)):
        if arr2[j] < arr2[cur_min_idx]:
            cur_min_idx = j
    
    arr2[i], arr2[cur_min_idx] = arr2[cur_min_idx], arr2[i]
print(arr2)
    






