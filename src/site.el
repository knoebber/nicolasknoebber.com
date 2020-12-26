;;; Package --- Summary
;;; Commentary:
;;
;; Set org config for publishing this website.

;;; Code:

(defconst html-main-head "<link rel=\"stylesheet\" type=\"text/css\" href=\"style.css\" />")
(defconst html-posts-head "<link rel=\"stylesheet\" type=\"text/css\" href=\"../style.css\" />")

;; (defun postamble-text (text)
;;   "Wraps TEXT in a span with class postamble-text."
;;   (format "<span class=\"postamble-text\">%s</span>" text))

;; (defun postamble-version (version)
;;   "Wraps VERSION in a span with class version-number."
;;   (format "<span class=\"postamble-text version-number\">%s</span>" version))

(defun publish-nicolasknoebber-file ()
  "Exports current org file to html and uploads to s3://nicolasknoebber.com."
  (interactive)
  (org-publish-current-file)
  (let* (
	 (org-file (buffer-file-name (buffer-base-buffer)))
	 (publishing-dir (org-publish-property :publishing-directory
                                               (org-publish-get-project-from-filename org-file)))
         (html-file (replace-regexp-in-string "org$" "html" (buffer-name)))
	 (html-file-path (concat publishing-dir "/" html-file))
	 (site-path (replace-regexp-in-string ".+personal-website" "" html-file-path))
         (aws-s3-cmd
          (concat "aws s3 cp " html-file-path " s3://nicolasknoebber.com" site-path)))
    (eshell-command aws-s3-cmd)))

(defconst html-postamble
  (concat
   "<span id=\"made-with\">
   exported by&nbsp;&nbsp;
   <a href=\"https://www.gnu.org/software/emacs\"
     ><img src=\"../logo/emacs.svg\" id=\"emacs-logo\" alt=\"Emacs\">
   </a>"
   emacs-version "&nbsp
   <a href=\"https://orgmode.org\">
     <img src=\"../logo/org-mode.svg\" id=\"org-mode-logo\" alt=\"Org\"></a>"
   org-version
   (format " on %s" (format-time-string "%m/%d/%y"))
   "</span>"))


(defconst html-posts-postamble
  (concat
   html-postamble
   "
<noscript>
  <div id=\"no-script-comment-message\">Enable scripts to see and post comments.</div>
</noscript>
<script type=\"text/javascript\" src=\"js/comments.js\"></script>"))

(defconst html-preamble "<a href=\"/\">Home</a>")
(defconst html-posts-preamble
  (concat html-preamble "<a href=\"/posts/\">Posts</a>"))

  
(defun generate-sitemap(title list)
  "Default site map, as a string.
TITLE is the title of the site map.  LIST is an internal
representation for the files to include, as returned by
`org-list-to-lisp'.  PROJECT is the current project.  This is
almost identical to the version in the org publish source code.
The only change I made is wrapping it in the .sitemap div."
  (concat
   "#+TITLE: " title
   "\n\n"
   "#+begin_sitemap\n"
   (org-list-to-org list)
   "\n#+end_sitemap"))

(defun format-sitemap-entry (entry _style project)
  "Format ENTRY in PROJECT."
    (format "[[file:%s][%s]] =%s="
	    entry
	    (org-publish-find-title entry project)
	    (format-time-string "%m/%d/%Y" (org-publish-find-date entry project))))

(defun format-exported-timestamps(timestamp _backend _channel)
  "Remove <> from exported org TIMESTAMP."
  (print (replace-regexp-in-string "&[lg]t;" "" timestamp))
  (replace-regexp-in-string "&[lg]t;" "" timestamp)
)

(eval-after-load 'ox
  '(add-to-list
    'org-export-filter-timestamp-functions
    'format-exported-timestamps))

(setq org-publish-project-alist
      `(("personal-website"
         :components ("main" "posts"))
	("main"
	 :publishing-directory "~/projects/personal-website"
	 :base-directory "~/projects/personal-website/src"
	 :publishing-function org-html-publish-to-html
	 :section-numbers nil
	 :with-toc nil
	 :html-head ,html-main-head
	 :html-preamble ,html-preamble
	 :html-postamble ,html-postamble
	 :html-head-include-scripts nil
	 :html-head-include-default-style nil
         )
	("posts"
         :publishing-directory "~/projects/personal-website/posts"
	 :base-directory "~/projects/personal-website/src/posts"
	 :publishing-function org-html-publish-to-html
	 :html-head ,html-posts-head
         :html-head-include-scripts nil
         :html-head-include-default-style nil
	 :html-preamble ,html-posts-preamble
	 :html-postamble ,html-posts-postamble
	 :auto-sitemap t
         :sitemap-title "Blog"
         :sitemap-function generate-sitemap
         :sitemap-format-entry format-sitemap-entry
         :sitemap-style list
         :sitemap-sort-files anti-chronologically
         :sitemap-filename "index.org"
         )))

;;; site.el ends here

