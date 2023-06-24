# jzhou 2023-06-20  ITMD-513
print("%s" % "Credit Card Validator Program")

# if the card number is valid, this function returns true 
def isValid(number):
    if getSize(number) < 13 or getSize(number) > 16:
        return False
    
    if not prefixMatched(number, 4) and not prefixMatched(number, 5) and not prefixMatched(number, 6) and not prefixMatched(number, 37):
        return False
        
    sum_of_odds = sumOfOddPlace(number)
    sum_of_evens = sumOfDoubleEvenPlace(number)
    total_sum = sum_of_evens + sum_of_odds
    if total_sum % 10 != 0:
        return False
    return True

# obtain the result from Step 2 
def sumOfDoubleEvenPlace(number):
    number_str = str(number)
    digits = [int(n) for n in number_str]
    reversed_digits = digits[::-1]
    sum = 0
    for i in range(len(reversed_digits)):
        if i % 2 == 1:
            sum += getDigit(reversed_digits[i] * 2)
    return sum

# if variable number is a single digit, then return the number  
# otherwise, return number as the sum of the two digits
def getDigit(number):
    if number < 10:
        return number
    else:
        return number // 10 + number % 10 

# this function returns the sum of odd place digits in variable number
def sumOfOddPlace(number):
    number_str = str(number)
    digits = [int(n) for n in number_str]
    reversed_digits = digits[::-1]
    sum = 0
    for i in range(len(reversed_digits)):
        if i % 2 == 0:
            sum += reversed_digits[i]
    return sum

# if the digit d is a prefix for variable number, then return true
def prefixMatched(number, d):
    prefix = getPrefix(number, 1)
    if d >= 10:
        prefix = getPrefix(number, 2)
    return str(d) == str(prefix)

# this function returns the number of digits in variable d 
def getSize(d):
    return len(str(d))

# returns the first k number of digits from variable number but if the
# number of digits in variable number is less than k, return number 
def getPrefix(number, k):
    number_str = str(number)
    if (len(number_str) < k):
        return number
    else:
        return int(number_str[:k])


card_numbers = [123456, 11111111111111, 4388576018452626, 4111111111111111]

for card_number in card_numbers:
    if isValid(card_number):
        print("Credit card number is valid. Card number: %d" % card_number)
    else:
        print("Credit card number is invalid. Card number: %d" % card_number)
