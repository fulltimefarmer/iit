# local variable declarations
totalCost = 0.0    # declare variable as a float type to accumulate total charges
appName = ""    # declare a variable for the appliance name
applianceList = []    # declare a variable for containing the list of appliances
costPerKW = 0.0    # declare a variable for the cost per KW - hr
annualUsage = 0.0    # declare a variable for the annual usage
costItems = 6    # declare a variable for number of cost items
totalKwhr = 0    # declare a variable for totalKwhr
variancePerKW = 0.0000    # declare a variable for variance

print ("[ please enter the requested data ]")

for x in range(costItems, 0, -1):
    appName = input("Appliance name %.d: " % (costItems + 1 - x))
    costPerKW = float(input("The cost per KW - hr of the appliance (in cents):"))
    annualUsage = float(input("The annual usage (in KW - hr):"))
    totalKwhr += costPerKW
    totalCost += costPerKW * annualUsage
    applianceList.append([appName, costPerKW, annualUsage])
    print("\n")
	
print("%-10s: $%.2f" % ("Total Annual Cost", totalCost))

averageKwhr = totalKwhr / costItems
print("%-10s: $%.2f" % ("Average", averageKwhr))

for i in applianceList:
    variancePerKW += (averageKwhr - i[1]) ** 2
    
variance = variancePerKW / costItems
print("%-10s: $%.2f" % ("Variance", variance))

print("%-10s: $%.2f" % ("stdDev", variance ** .5))
