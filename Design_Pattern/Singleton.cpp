
//
// Created by lidaomeng on 2018/4/26.
//

#include <mutex> // 使用C++11互斥锁
#include <iostream>
#include <string>
using namespace std;

std::mutex g_Mutex;

class Person{
private:
    string m_Name;
    unsigned m_Age;
public:
    Person():m_Age(0){}
    Person(const string& aoName, unsigned aiAge):m_Name(aoName), m_Age(aiAge){}

public:
    void SetPerson(const string& aoName, unsigned aiAge){
        m_Name = aoName;
        m_Age = aiAge;
    }
    void Print_Person(){
        cout << "Name: " << m_Name << ", Age: " << m_Age << endl;
    }
};

// 版本1 类模板不带参数
template <class T>
class Singleton{
protected:
    Singleton(){}
    ~Singleton(){}
private:
    Singleton(const Singleton&);
    Singleton& operator=(const Singleton&);

    //static std::mutex m_Mutex;
    static T* m_pInstance;
public:
    static T* GetInstance();
    static void DeleteInstance();
};

template <class T>
T* Singleton<T>::m_pInstance = nullptr;

template <class T>
T* Singleton<T>::GetInstance() {
    if(m_pInstance == nullptr){
        g_Mutex.lock();
        if(m_pInstance == nullptr){ // Note: 双判断
            T* pTemp = new T(); // Note: 临时指针
            m_pInstance = pTemp;
        }
        g_Mutex.unlock();
    }
    return m_pInstance;
}

template <class T>
void Singleton<T>::DeleteInstance() {
    if(m_pInstance != nullptr){
        g_Mutex.lock();
        if(m_pInstance != nullptr){ // Note: 再次判断指针是否为空，否则会出现释放空指针现象
            delete m_pInstance;
            m_pInstance = nullptr;
        }
        g_Mutex.unlock();
    }
}

int main(){
    Singleton<Person>::GetInstance()->Print_Person();

    Singleton<Person>::GetInstance()->SetPerson("li dao meng", 99);
    Singleton<Person>::GetInstance()->Print_Person();

    return 0;
}
