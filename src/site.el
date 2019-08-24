;;; Package --- Summary
;;; Commentary:
;;
;; Set org config for publishing this website.
;;
;; Site Map
;; TODO use tree.

;;; Code:
(setq org-publish-project-alist
      '(("personal-website"
	 :base-directory "~/projects/personal-website/src"
	 :publishing-directory "~/projects/personal-website"
	 :publishing-function org-html-publish-to-html
	 :section-numbers nil
	 :with-toc nil
	 :with-title nil
	 :html-head "<link rel=\"stylesheet\" type=\"text/css\" href=\"style.css\" />"
	 :html-preamble "<a href=\"/\">Nicolas Knoebber</a>"
	 :html-postamble nil
	 :html-head-include-scripts nil
	 :html-head-include-default-style nil
	 )))

;;; site.el ends here
