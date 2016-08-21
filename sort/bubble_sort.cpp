#include "extruct.cpp"

#include <iostream>

using namespace std;

class bubble_sort{

private:
  static void change(node * n1, node *n2){
    float *a = n1->get_info();
    n1->set_info(n2->get_info());
    n2->set_info(a);
  }
public:
  static void sort(node * r){
    node * i = r;
    node *j=NULL;
    float vi=0;
    float vj=0;
    while(i->next!=NULL){
      j=i->next;
      while(j->next!=NULL){
        vi = i->get_value();
        vj = j->get_value();
        if(vi>vj) change(i,j);
        j=j->next;
      }
      i=i->next;
    }
  }


};

int main(){
  node * n = new node(-1);
  n->append(1);
  n->append(5);
  n->append(4);
  n->append(2);
  n->append(5);

  n->print_all();
cout<<"\nSorted:"<<endl;
bubble_sort::sort(n);
n->print_all();


  return 0;
}
