# RLHF

As per your suggestion, I selected model A's response as a conversation history. Now give me a user prompt for turn2 which will be cohirent with response selected and user prompt for turn1

Rework Done: Added 2 turns with different prompts

Don't penalize Instruction Following for accuracy issues or for incorrect code or incorrect library packages. Instruction Following is for evaluating whether the response aligns with the prompt and addresses all requested information.

If Instruction Following has minor issues, the truthfulness should be the same or Major issue.
Do not penalize IF for accuracy issues, penalize Truthfulness instead

Now provide me these details for both the responses 

1. Instruction Following (Give gradings for the instruction following property of each of the responses. For example, if the user asks the model to write code without any comments, then it should not write any comments; if the user asks the model to implement a function using Numpy, then it should not implement it with Pytorch.Rate it out off this three 1. No issues 2. Minor issues 3. Major issues )
2. Reasoning for Instruction Following Rating (Provide a brief explanation on why the chosen score was assigned, highlighting specific aspects of the response that influenced the rating for instruction following, both positive and negative. Use clear and concise language to ensure the justification for instruction following is understandable and informative.)
3. Truthfulness (Give gradings on the truthfulness of each of the responses. Here, truthfulness means the correctness of the response. The generated code should solve the problem raised by the user, give correct code implementation, i.e., no logical and syntax errors, etc. use this values 1. No issues 2. Minor issues 3. Major issues )
4. Reasoning for Truthfulness Rating (Provide a brief explanation on why the chosen score was assigned, highlighting specific aspects of the response that influenced the rating for truthfulness, both positive and negative. Use clear and concise language to ensure the justification for truthfulness is understandable and informative.)
5. Conciseness (Rate the verbosity of each of the responses. use this values 1. just right 2. too verbose 3. too short)
6. Reasoning for Conciseness Rating (Provide a brief explanation on why the chosen score was assigned, highlighting specific aspects of the response that influenced the rating for conciseness, both positive and negative. Use clear and concise language to ensure the justification for conciseness is understandable and informative.)
7. Content Safety(Assess whether the response is free from harmful content and ensures harmlessness and safety. use this values 1. 1. No issues 2. Minor issues 3. Major issues)
8. Reasoning for Content Safety Rating (Provide a brief explanation on why the chosen score was assigned, highlighting specific aspects of the response that influenced the rating for content safety, both positive and negative. Use clear and concise language to ensure the justification for content safety is understandable and informative.)
9. Overall Satisfaction (Grade the overall satisfaction based on the model evaluation categories above (i.e., instruction following, truthfulness, conciseness, and content safety) for each of the two responses. Give the grading using the following values: 1. Amazing, 2. Pretty good, 3. OK, 4. Pretty Bad, 5. Horrible ) 


10. Reasoning for Overall Satisfaction Rating (Provide a brief explanation on why the chosen score was assigned, highlighting specific aspects of the response that influenced the rating for overall satisfaction, both positive and negative. Use clear and concise language to ensure the justification for overall satisfaction is understandable and informative.)



A. Preference Explanation.

B. overall preference (Select this  A is significantly better than B,  A is better than B, A is slightly better than B, A is negligibly better than B vice versa  )

1. 52
2. 72
3. 64

Mon --> 14/15
390816 --> 2
390901 --> 2 
390723 --> 2
390835 --> 2
390843 --> 2
390871 --> 2
390864 --> 2

Tuesday --> 14/15
390865 --> 2
390834 --> 2
391019 --> 1+2
391041 --> 1+2
390954 --> 2
390XXX --> 2

While evaluating do deep analysis also refer AI model training guidelines you have in memory
Wed -->  20/15
390962 --> 2+2
390654 --> 2+2
390814 --> 2
390844 --> 2
390886 --> 2
464832 --> 2
464811 --> 2+2

Thu 16/15
464839 --> 2+2
464847 --> 2
464929 --> 2
464926 --> 2
464975 --> 2
464938 --> 2+2

Fri 13/15
464990 --> 2
465005 --> 2
465081 --> 2 + 2
465074 --> 3
464914 --> 2

Mon
464993 --> 
Rework 
2 + 2 