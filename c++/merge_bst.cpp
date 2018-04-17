#include <vector>

struct TreeNode {
	int val;
	TreeNode *left;
	TreeNode *right;
	TreeNode(int x) : val(x), left(NULL), right(NULL) {}
};
 
class Solution {
public:
    TreeNode* mergeTrees(TreeNode* t1, TreeNode* t2) {

        if ((t1 != NULL) and (t2 != NULL)) {
        	t1->val = t1->val + t2->val; 
        } else {
        	if (t2 != NULL) {
        		t1 = t2;
        	}
        }
        if ((t1 != NULL) and (t1->left != NULL) ){
        	if ((t2 != NULL) and t2->left != NULL) {
        		mergeTrees(t1->left, t2->left);
        	}
        } else {
        	if ((t2 != NULL) and t2->left != NULL) {
        		t1->left = t2->left;
        	}
        }

        if ((t1 != NULL) and t1->right != NULL) {
        	if ((t2 != NULL) and t2->right != NULL) {
        		mergeTrees(t1->right, t2->right);
        	}
        } else {
        	if ((t2 != NULL) and t2->right != NULL) {
        		t1->right = t2->right;
        	}
        }
        return t1;
    }
};