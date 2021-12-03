#include <iostream>
#include <stdio.h>
#include <assert.h>
#include <inttypes.h>
#include <string.h>
#include <x86intrin.h>
#include <math.h>
#include <unordered_map>
#include <algorithm>
#include <bitset>
#include <vector>
#include <fstream>
#include <chrono>
#include <iomanip>

using namespace std;

int main(int argc, char** argv){

    auto now = std::chrono::system_clock::now();
    auto now_c = std::chrono::system_clock::to_time_t(now);
    std::chrono::steady_clock::time_point begin = std::chrono::steady_clock::now();

    vector <unsigned short int> input;
    string line;    
    ifstream fin("input");
    while(getline(fin,line)){
        input.push_back(stoi(line));
    }

    int total;
    for (int x = 0; x < input.size(); x+=15){
        unsigned short int a[16];
        memmove(a, &input[x], sizeof a);

        unsigned short int b[16];
        memmove(b,  a + 1, sizeof a);
        b[15] = b[14];

        __m256i v0 = _mm256_loadu_si256((__m256i*)(a));
        __m256i v1 = _mm256_loadu_si256((__m256i*)(b));

        __m256i v2 = _mm256_cmpgt_epi16 (v1, v0);
        __m256i v3 = _mm256_abs_epi16(v2);
        __m128i v4 = _mm256_extractf128_si256(v3, 0);
        __m128i v5 = _mm256_extractf128_si256(v3, 1);
        __m128i v6 = _mm_add_epi16(v4, v5);

        ushort sum[8];
        _mm_storeu_si128((__m128i*)&sum[0], v6);

        total = total + (sum[0]+sum[1]+sum[2]+sum[3]+sum[4]+sum[5]+sum[6]+sum[7]);

    }
    std::cout << total << std::endl;


    std::chrono::steady_clock::time_point end = std::chrono::steady_clock::now();
    double seconds = (double)std::chrono::duration_cast<std::chrono::microseconds> (end - begin).count() / 1000000;

    std::cout << seconds  << endl;
    std::cout << std::endl;

    return 0;
}

