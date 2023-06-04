/*******
* @Author:qingmeng
* @Description:
* @File:check
* @Date2022/2/16
 */

package tool

// CheckIfSensitive 检查敏感词
func CheckIfSensitive(s string) bool {
	sensitiveWords = append(sensitiveWords, "你妈", "傻逼", "反中国", "反人类")
	trie := NewTrie()
	trie.Insert(sensitiveWords)
	bt := []interface{}{s}
	return trie.StartsWith(bt)
}
