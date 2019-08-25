;;; Package --- Summary
;;; Commentary:
;;
;; Set org config for publishing this website.
;;
;; Site Map
;; TODO use tree.

;;; Code:

(defconst html-style "<link rel=\"stylesheet\" type=\"text/css\" href=\"style.css\" />")

(defconst html-main-head
  (concat "<link rel=\"stylesheet\" type=\"text/css\" href=\"includes.css\" />" html-style))

(defconst html-posts-head
  (concat "<link rel=\"stylesheet\" type=\"text/css\" href=\"../includes.css\" />" html-style))

(defun postamble-text (text)
  "Wraps TEXT in a span with class postamble-text."
  (format "<span class=\"postamble-text\">%s</span>" text))

(defun postamble-version (version)
  "Wraps VERSION in a span with class version-number."
  (format "<span class=\"postamble-text version-number\">v%s</span>" version))

(defconst html-postamble
  (concat
   "<span id=\"made-with\">"
   (postamble-text "powered by&nbsp;")
   "<a href=\"https://www.gnu.org/software/emacs\">"
   "<img src=\"../logo/emacs.svg\" id=\"emacs-logo\" alt=\"emacs logo\">"
   "</a>"
   (postamble-version emacs-version)
   "&nbsp<a href=\"https://orgmode.org\">"
   "<img src=\"../logo/org-mode.svg\" id=\"org-mode-logo\" alt=\"org mode logo\">"
   "</a>"
   (postamble-version org-version)
   "</span>"
   "<span id=\"copyright\">"
   (format "© %s" (format-time-string "%Y"))
   "</span>"))


(setq org-publish-project-alist
      `(("personal-website"
         :components ("main" "posts"))
	("main"
	 :publishing-directory "~/projects/personal-website"
	 :base-directory "~/projects/personal-website/src"
	 :publishing-function org-html-publish-to-html
	 :section-numbers nil
	 :with-toc nil
	 :with-title nil
	 :html-head ,html-main-head
	 :html-preamble "<a href=\"/\">Nicolas Knoebber</a>"
	 :html-postamble ,html-postamble
	 :html-head-include-scripts nil
	 :html-head-include-default-style nil)
	("posts"
         :publishing-directory "~/projects/personal-website/posts"
	 :base-directory "~/projects/personal-website/src/posts"
	 :publishing-function org-html-publish-to-html
	 :section-numbers nil
	 :with-toc nil
	 :html-head ,html-posts-head
	 :html-head-include-scripts nil
	 :html-head-include-default-style nil
	 :html-preamble "<a href=\"../blog.html\">Blog</a>"
	 :html-postamble ,html-postamble
	 )))


;;; site.el ends here
