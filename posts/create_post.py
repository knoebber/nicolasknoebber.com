#!/usr/bin/python3
from markdown import markdown
from sys      import argv

"""
adds a new <tr> element to blog.html
new row will always be the first to keep reverse chronoligical order
"""
def add_entry_to_list(post_num) :
  post_num = str(post_num)
  md     = open(post_num+'.md')
  header = md.readline()[3:-1] #slice omits the beginning hashes and trailing \n
  date   = md.readline()[5:-1]
  md.close()
  html = open('../blog.html')
  lines = html.readlines()
  html.close()
  new_element = '<tr><td><a href="posts/'+post_num+'.html">'+header+'</a></td><td>'+date+'</td></tr>\n'
  i = len(lines) - 1
  while i > 0 :
    if lines[i].strip() == new_element.strip() :
      print('list item already exists')
      return
    if lines[i].strip() == '<tbody>' :
      lines.insert(i+1,' '*4+new_element) #indent new tag properly and add to file
      print('new list item created')
      html = open('../blog.html','wt')
      html.writelines(lines)
      html.close()
      return
    i -= 1

"""
reads from the markdown file (post_num).md and writes (post_num).html
"""
def add_entry(post_num) :
  try :
    md = open(post_num+'.md')
  except :
    print(post_num+'.md'+' does not exist!')
    return
  h = open('header.html')
  header = h.read()
  h.close()
  html = markdown(md.read())
  html = '<LINK REL=StyleSheet HREF="style.css" TYPE="text/css">\n' + html
  html = header + html
  md.close()
  post = open(post_num+'.html','wt')
  post.write(html)
  post.close()
  print('post updated')
  add_entry_to_list(post_num)

if __name__ == '__main__' :
  post_num = argv[1]
  add_entry(post_num)
