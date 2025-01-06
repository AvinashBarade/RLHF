# RLHF

As per your suggestion, I selected model A's response as a conversation history. Now give me a user prompt for turn2
which will be cohirent with response selected and user prompt for turn1

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

1. 47


Monday 12/15
494133 --> 2
494156 --> 2C
494158 --> 4IRC
494157 --> 2
494151 --> 2

Tue 15/15
494242 --> 2
494240 --> 4IR
494235 --> 3
494231 --> 2
494212 --> 4IR

Wed 06/15
494055 --> 2
494021 --> 2
494015 --> 2

Friday 14/15
494085 --> 2
494187 --> 2
494142 --> 2
493944 --> 2
494106 --> 2
494198 --> 2
494181 --> 2