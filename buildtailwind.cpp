#include <iostream>
#include <cstdlib>

int main(){
    const char* command = "npx tailwindcss -i ./input.css -o ./ui/static/css/output.css";

    int result = system(command);

    if(result == 0){
        std::cout << "Command executed successfully" <<
        std::endl;
    } else{
        std::cerr << "Error executing command" << std::endl;
    }
    return 0;
}
