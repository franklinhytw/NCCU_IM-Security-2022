// NCCU Information Security HW1 - Vigenère Cipher Encoder/Decoder
// Feb 23, 2022
// 109971014 林翰陽

#include <iostream>
#include <limits>
#include <string>
#include <vector>
#include <map>
#include <cctype>
#include <cstring>

#ifdef _WIN32
#include <windows.h>
#endif

using namespace std;

void inputKey(string &key) {
  cout << "Please Input the KEY(0~9 A~Z):";
  cin >> key;

  if (key.empty()) {
    cout << "NO KEY WAS FOUND" << endl;
    return;
  }
}

void inputText(string &text, bool is_plain_text) {
  if (is_plain_text) {
    cout << "Please Input the PLAIN TEXT(0~9 A~Z):";
    cin >> text;

    if (text.empty()) {
      cout << "NO PLAIN TEXT WAS FOUND" << endl;
      return;
    }
  } else {
    cout << "Please Input the CYPHER TEXT:";
    cin >> text;

    if (text.empty()) {
      cout << "NO CYPHER TEXT WAS FOUND" << endl;
      return;
    }
  }
}

vector<uint8_t> encode_to_numeric(string str) {
    vector<uint8_t> encode_text;
    for(auto &c : str) {
        // a -> A, z -> Z
        int upper_char = toupper(c);

        // A~Z
        if(isalpha(upper_char)) {
            encode_text.push_back((uint8_t)(upper_char - 55));
        }
        // 0~9
        else if(isdigit(upper_char)) {
            encode_text.push_back((uint8_t)(upper_char - 48));
        }
    }
    return encode_text;
}

string decode_to_string(vector<uint8_t> num_list) {
    string text;
    
    for(auto &num : num_list) {
        // cout << (int)num << ", ";
        // 0~9
        if(num < 10) {
            text.push_back(num+48);
        }
        // A~Z
        else {
            text.push_back(num+55);
        }
    }

    return text;
}

string encrypt(string key, string plaintext) {
    vector<uint8_t> encoded_cipher_list;

    auto encoded_key = encode_to_numeric(key);
    auto encoded_text = encode_to_numeric(plaintext);

    size_t key_length = encoded_key.size();
    size_t key_idx_pointer = 0;

    for(auto &v : encoded_text) {
        // printf("[%u] enc=%u, v=%u, ans=%d\n", key_idx_pointer, encoded_key[key_idx_pointer], v, (int)(v + encoded_key[key_idx_pointer]) % 36);
        encoded_cipher_list.push_back((v + encoded_key[key_idx_pointer]) % 36);

        if(key_idx_pointer+1 == key_length) key_idx_pointer = 0;
        else key_idx_pointer++;
    }

    return decode_to_string(encoded_cipher_list);
}

string decrypt(string key, string ciphertext) { 
    vector<uint8_t> encoded_cipher_list;

    auto encoded_key = encode_to_numeric(key);
    auto encoded_text = encode_to_numeric(ciphertext);

    size_t key_length = encoded_key.size();
    size_t key_idx_pointer = 0;

    for(auto &v : encoded_text) {
        // printf("[%u] enc=%u, v=%u, ans=%d\n", key_idx_pointer, encoded_key[key_idx_pointer], v, (int)(36 + (v - encoded_key[key_idx_pointer]) % 36) % 36);
        encoded_cipher_list.push_back((36 + (v - encoded_key[key_idx_pointer]) % 36) % 36);

        if(key_idx_pointer+1 == key_length) key_idx_pointer = 0;
        else key_idx_pointer++;
    }

    return decode_to_string(encoded_cipher_list);
}

int main(int argc, char const *argv[]) {
#ifdef _WIN32
  SetConsoleOutputCP(65001);
#endif

  while (1) {
    cout << "[Vigenère Cipher Encoder/Decoder]\n" << endl;
    cout << "SELECT function:" << endl;
    cout << "1 -> Encryption" << endl;
    cout << "2 -> Decryption" << endl;
    cout << "0 -> QUIT program" << endl;

    int x;

    while (!(cin >> x)) {
      cin.clear();
      cin.ignore(numeric_limits<streamsize>::max(), '\n');
      cout << "Invalid input. Try again: ";
    }

    switch (x) {
      case 1: {
        cout << "[Encryption]" << endl;
        string key, plaintext;

        inputKey(key);
        if (key.empty()) break;

        inputText(plaintext, true);
        if (plaintext.empty()) break;

        printf("The cipher text is >>>[%s]<<<", encrypt(key, plaintext).c_str());
        break;
      }

      case 2: {
        cout << "[Decryption]" << endl;
        string key, ciphertext;

        inputKey(key);
        if (key.empty()) break;

        inputText(ciphertext, false);
        if (ciphertext.empty()) break;

        printf("The plain text is >>>[%s]<<<", decrypt(key, ciphertext).c_str());
        break;
      }

      case 0: {
        exit(0); // EXIT PROGRAM
      }
    }

    cout << "\n\n";
  }

  return 0;
}