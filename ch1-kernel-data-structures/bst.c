#include <stdio.h>
#include <stdlib.h>

typedef struct Node {
    int data;
    struct Node *left;
    struct Node *right;
} Node;

Node* createNode(int data){
    Node* newNode = (Node*) malloc(sizeof(Node));
    newNode->data = data;
    newNode->right = NULL;
    newNode->left = NULL;

    return newNode;
}

Node* insert(Node* root, int data){
    if (root == NULL) return createNode(data);

    if (data < root->data)
        root->left = insert(root->left, data);
    else
        root->right = insert(root->right, data);

    return root;
}

Node* search(Node* root, int key){
    if (root == NULL || root->data == key) return root;

    if (key > root->data) return search(root->right, key);

    return search(root->left, key);
}

void inOrderTraversal(Node* root){
    if (root != NULL){
        inOrderTraversal(root->left);
        printf("%d ", root->data);
        inOrderTraversal(root->right);
    }
}

void preOrderTraversal(Node* root){
    if (root != NULL){
        printf("%d ", root->data);
        preOrderTraversal(root->left);
        preOrderTraversal(root->right);
    }
}

void postOrderTraversal(Node* root){
    if (root != NULL){
        postOrderTraversal(root->left);
        postOrderTraversal(root->right);
        printf("%d ", root->data);
    }
}



int main()
{
    // Soon, BST implementation will be written
    Node* root = NULL;
    root = insert(root, 47);
    insert(root, 30);
    insert(root, 20);
    insert(root, 33);
    insert(root, 54);
    insert(root, 51);
    insert(root, 52);

    printf("Inorder traversal of the BST: ");
    inOrderTraversal(root);

    printf("\n\n");
    printf("PreOrder traversal of the BST: ");
    preOrderTraversal(root);

    printf("\n\n");
    printf("PostOrder traversal of the BST");
    postOrderTraversal(root);

    return 0;
}
