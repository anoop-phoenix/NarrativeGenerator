# Narrative Generator
* Generate narrative (story) when executing a sequence of statement. 
* If there is a predicate and it returns `False`, the process will be terminated.
* The narrative can be represented as a paragraph with indent levels or a json like string. 

## The sample problem
### Sample 1 -> Evaluate() function
The process checks wether a student is a good student or not. To be qualified for that hornor, the guy must have:
* All majors' overall are greater or equal 8.5
* With each major, there is no assignment's mark less than 5.0
* Certificate of IELTS

If the student has the certificate of JLPT, he or she is qualified for a very good student hornor.

### Sample 2 -> Average() function
The process also return the overall score of the student. Show reference to overal score of each subject.

**Beside**, the algorithm must provide the story evaluating the student.

## Note
* Class `EvaluationProcesses` has its own print method `GenerateStory` 


## Experience
* Only provide `ReferenceResource` in the leaf node, therefore, the information is not duplicated 
* The detail of `narrative` depends on the intent of the author