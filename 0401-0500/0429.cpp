class Solution {
public:
    vector<vector<int>> levelOrder(Node* root) {
        // 特殊情况 ： root 为空
        if (root == nullptr) {
            return {};
        }

        vector<vector<int>> res;
        queue<Node *> q;
        
        q.push(root);

        while (!q.empty())
        {
            int size = q.size();
            vector<int> level;
            for (int i = 0; i < size; i++)
            {
                Node *node = q.front();
                q.pop();

                level.push_back(node->val);

                for(auto child : node->children)
                {
                    q.push(child);
                }
            }
            res.push_back(level);
            //res.push_back(move(level));
        }
        return res;
    }
};


