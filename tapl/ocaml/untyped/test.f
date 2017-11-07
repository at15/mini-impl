/* Examples for testing */

x/;
x;

lambda x. x;
(lambda x. x) (lambda x. x x); 

/* added by @at15 for https://github.com/at15/reika/issues/7 */
lambda x. x x;
(lambda x. x) (lambda y. y);
