class Solution {
public:
    vector<vector<int>> levelOrder(TreeNode* root) {
        // 特殊情况：root 空
        if(root == nullptr) return {};

        vector<vector<int>> res;

        queue<TreeNode*> q;
        q.push(root);

        while (!q.empty()) {

            int level_size = q.size();
            vector<int> level;

            for (int i = 0; i < level_size; i++) {
                TreeNode* node = q.front();
                q.pop();

                level.push_back(node->val);

                if (node->left) q.push(node->left);
                if (node->right) q.push(node->right);
            }
            res.push_back(move(level));
        }
        return res;
    }
};