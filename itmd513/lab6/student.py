class Student:

    Scores = {}

    # initializing the constructor method

    def __init__(self, name, grade):
        self.name = name
        self.grade = grade

    def getScores(self):

        answer_key = []
        # read into answer_key list, the answer key from file
        answer_key = [line.strip() for line in open("answers.txt", 'r')]

        student_answers = []
        # read into student_answers list, student answers from file
        student_answers = [line.strip().split(',') 
                           for line in open("data.txt", 'r')]
        total_score = 100

        # place additional code statements here for the above function

        #---start your loop processing logic here---#

        #---end your loop processing logic here---#
        
        #---continue the class definition#
        
        Student.Scores[self.getName()] = total_score

    def getName(self):
        return self.name

    @staticmethod
    def sortDict():
        return sorted(Student.Scores.items())

#---end the class definition#

student_objs = [
    Student('Sammy Student', 65),
    Student('Betty sanchez', 45),
    Student('Alice brown', 100),
    Student('tom Schulz', 50),
]

for index in range(len(student_objs)):
    student_objs[index].getScores()

sortList = Student.sortDict()

for k, v in sortList:
    print(k, "has score:", v)