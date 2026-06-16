import os
import re

def main():
    # Get the directory of the script
    repo_dir = os.path.dirname(os.path.abspath(__file__))
    
    # 1. Scan the repository for subdirectories matching \d{4}
    # Using regex to match exactly 4 digits
    folder_pattern = re.compile(r'^\d{4}$')
    solved_count = 0
    for name in os.listdir(repo_dir):
        if os.path.isdir(os.path.join(repo_dir, name)) and folder_pattern.match(name):
            solved_count += 1
            
    # 2. Update/create README.md
    readme_path = os.path.join(repo_dir, 'README.md')
    
    # Read the existing content of README.md (if it exists)
    if os.path.exists(readme_path):
        with open(readme_path, 'r', encoding='utf-8') as f:
            content = f.read()
    else:
        content = "# leet-code\n\n"
        
    start_marker = "<!-- SUMMARY:START -->"
    end_marker = "<!-- SUMMARY:END -->"
    
    summary_content = f"{start_marker}[![Solved Challenges](https://img.shields.io/badge/Solved%20Challenges-{solved_count}-brightgreen?style=for-the-badge&logo=leetcode&logoColor=white)](https://leetcode.com/){end_marker}"

    # If the markers aren't in the file, append them
    if start_marker not in content or end_marker not in content:
        # If content doesn't end with a newline, add one
        if content and not content.endswith('\n'):
            content += '\n'
        content += '\n' + summary_content + '\n'
    else:
        # Replace content between markers
        pattern = re.compile(rf"{re.escape(start_marker)}.*?{re.escape(end_marker)}", re.DOTALL)
        content = pattern.sub(summary_content, content)
        
    # Write back to README.md
    with open(readme_path, 'w', encoding='utf-8') as f:
        f.write(content)
        
    print(f"Updated README.md: {solved_count} problems solved")

if __name__ == '__main__':
    main()
