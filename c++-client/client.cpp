#include <string>
#include <iostream>
#include <curl/curl.h>

using namespace std;

static size_t call_back(char* ptr, size_t size, size_t nmemb, std::string* str) {
  int realsize = size * nmemb;
  str->append(ptr, realsize);
  return realsize;
}
static std::string curl_post_wrapper(std::string url, std::string post_data) {
  CURL *curl;
  CURLcode res;
  curl = curl_easy_init();
  if( curl == nullptr ) {
    return "Error";
  }
  string chunk = "";
  curl_easy_setopt(curl, CURLOPT_URL, url.c_str());
  curl_easy_setopt(curl, CURLOPT_POSTFIELDS, post_data.c_str());
  curl_easy_setopt(curl, CURLOPT_POSTFIELDSIZE, post_data.size());
  curl_easy_setopt(curl, CURLOPT_WRITEFUNCTION, call_back);
  curl_easy_setopt(curl, CURLOPT_WRITEDATA, static_cast<string*>(&chunk));
  res = curl_easy_perform(curl);
  curl_easy_cleanup(curl);
  return chunk;
}
int main() {
  std::string res = curl_post_wrapper("http://localhost:4567/test", "{\"type\":\"ADD\",\"data\":4.0}");
  cout << res << endl;
}
