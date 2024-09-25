<p>书店店员有一张链表形式的书单，每个节点代表一本书，节点中的值表示书的编号。为更方便整理书架，店员需要将书单倒过来排列，就可以从最后一本书开始整理，逐一将书放回到书架上。请倒序返回这个书单链表。</p>

<p>&nbsp;</p>

<p><strong>示例 1：</strong></p>

<pre>
<strong>输入：</strong>head = [3,6,4,1]

<strong>输出：</strong>[1,4,6,3]
</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<p><code>0 &lt;= 链表长度 &lt;= 10000</code></p>

<details><summary><strong>Related Topics</strong></summary>栈 | 递归 | 链表 | 双指针</details><br>

<div>👍 480, 👎 0<span style='float: right;'><span style='color: gray;'><a href='https://github.com/labuladong/fucking-algorithm/discussions/939' target='_blank' style='color: lightgray;text-decoration: underline;'>bug 反馈</a> | <a href='https://labuladong.online/algo/fname.html?fname=jb插件简介' target='_blank' style='color: lightgray;text-decoration: underline;'>使用指南</a> | <a href='https://labuladong.online/algo/images/others/%E5%85%A8%E5%AE%B6%E6%A1%B6.jpg' target='_blank' style='color: lightgray;text-decoration: underline;'>更多配套插件</a></span></span></div>

<div id="labuladong"><hr>

**通知：已完成网站教程、网站习题、配套插件中所有多语言代码的校准，解决了之前 chatGPT 翻译可能出错的问题~**

<details><summary><strong>labuladong 思路</strong></summary>

<div id="labuladong_solution_zh">

## 基本思路

这题解法很多，比如我们可以借助「栈」这种先进后出的结构来得到链表的倒序遍历结果，但是本文使用 [东哥手把手带你刷二叉树（纲领篇）](https://labuladong.online/algo/essential-technique/binary-tree-summary/) 讲到的后序遍历技巧来写代码。

递归函数本质上就是让操作系统帮我们维护了递归栈，和栈的解法效率上完全相同，但是这样写代码有助于我们深入理解递归的思维。

</div>

**标签：[栈](https://labuladong.online/algo/)，[链表](https://labuladong.online/algo/)**

<div id="solution">

## 解法代码



<div class="tab-panel"><div class="tab-nav">
<button data-tab-item="cpp" class="tab-nav-button btn " data-tab-group="default" onclick="switchTab(this)">cpp🤖</button>

<button data-tab-item="python" class="tab-nav-button btn " data-tab-group="default" onclick="switchTab(this)">python🤖</button>

<button data-tab-item="java" class="tab-nav-button btn active" data-tab-group="default" onclick="switchTab(this)">java🟢</button>

<button data-tab-item="go" class="tab-nav-button btn " data-tab-group="default" onclick="switchTab(this)">go🤖</button>

<button data-tab-item="javascript" class="tab-nav-button btn " data-tab-group="default" onclick="switchTab(this)">javascript🤖</button>
</div><div class="tab-content">
<div data-tab-item="cpp" class="tab-item " data-tab-group="default"><div class="highlight">

```cpp
// 注意：cpp 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
// 本代码已经通过力扣的测试用例，应该可直接成功提交。

class Solution {
public:
    vector<int> reversePrint(ListNode* head) {
        traverse(head);
        return res;
    }

    // 记录链表长度
    int len = 0;
    // 结果数组
    vector<int> res;
    // 结果数组中的指针
    int p = 0;

    // 递归遍历单链表
    void traverse(ListNode* head) {
        if (head == nullptr) {
            // 到达链表尾部，此时知道了链表的总长度
            // 创建结果数组
            res.resize(len);
            return;
        }
        len++;
        traverse(head->next);
        // 后序位置，可以倒序操作链表
        res[p] = head->val;
        p++;
    }

    // 用「分解问题」的思路写递归解法
    // 因为 C++ 的 vector 可以支持 push_back 操作，所以我们不需要改变返回值类型
    vector<int> reversePrint2(ListNode* head) {
        // base case
        if (head == nullptr) {
            return {};
        }

        // 把子链表翻转的结果算出来，示例 [3,2]
        vector<int> subProblem = reversePrint2(head->next);
        // 把 head 的值接到子链表的翻转结果的尾部，示例 [3,2,1]
        subProblem.push_back(head->val);
        return subProblem;
    }
};
```

</div></div>

<div data-tab-item="python" class="tab-item " data-tab-group="default"><div class="highlight">

```python
# 注意：python 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
# 本代码已经通过力扣的测试用例，应该可直接成功提交。

class Solution:
    def reversePrint(self, head: ListNode) -> List[int]:
        # 用「遍历」的思路写递归解法
        def traverse(head):
            nonlocal len_, res, p
            if not head:
                res = [0] * len_
                return
            len_ += 1
            traverse(head.next)
            res[p] = head.val
            p += 1
        
        len_ = 0
        res = []
        p = 0
        traverse(head)
        return res

    def reversePrint2(self, head: ListNode) -> List[int]:
        # 用「分解问题」的思路写递归解法
        def sub_problem(head):
            if not head:
                return []
            sub_res = sub_problem(head.next)
            sub_res.append(head.val)
            return sub_res
        
        return sub_problem(head)
```

</div></div>

<div data-tab-item="java" class="tab-item active" data-tab-group="default"><div class="highlight">

```java
// 用「遍历」的思路写递归解法
class Solution {
    public int[] reversePrint(ListNode head) {
        traverse(head);
        return res;
    }

    // 记录链表长度
    int len = 0;
    // 结果数组
    int[] res;
    // 结果数组中的指针
    int p = 0;

    // 递归遍历单链表
    void traverse(ListNode head) {
        if (head == null) {
            // 到达链表尾部，此时知道了链表的总长度
            // 创建结果数组
            res = new int[len];
            return;
        }
        len++;
        traverse(head.next);
        // 后序位置，可以倒序操作链表
        res[p] = head.val;
        p++;
    }

    // 用「分解问题」的思路写递归解法
    // 因为 Java 的 int[] 数组不支持 add 相关的操作，所以我们把返回值修改成 List
    // 定义：输入一个单链表，返回该链表翻转的值，示例 1->2->3
    List<Integer> reversePrint2(ListNode head) {
        // base case
        if (head == null) {
            return new LinkedList<>();
        }

        // 把子链表翻转的结果算出来，示例 [3,2]
        List<Integer> subProblem = reversePrint2(head.next);
        // 把 head 的值接到子链表的翻转结果的尾部，示例 [3,2,1]
        subProblem.add(head.val);
        return subProblem;
    }
}
```

</div></div>

<div data-tab-item="go" class="tab-item " data-tab-group="default"><div class="highlight">

```go
// 注意：go 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
// 本代码不保证正确性，仅供参考。如有疑惑，可以参照我写的 java 代码对比查看。

/*
用「遍历」的思路写递归解法
*/

func reversePrint(head *ListNode) []int {
    var res []int
    var len int

    // 递归遍历单链表
    var traverse func(*ListNode)
    traverse = func(head *ListNode) {
        if head == nil {
            // 到达链表尾部，此时知道了链表的总长度
            // 创建结果数组
            res = make([]int, len)
            return
        }
        len++
        traverse(head.Next)
        // 后序位置，可以倒序操作链表
        res[len-p-1] = head.Val
        p++
    }

    traverse(head)
    return res
}

/*
用「分解问题」的思路写递归解法
因为 Go 不支持泛型，所以我们把返回值修改成 []int
定义：输入一个单链表，返回该链表翻转的值，示例 1->2->3
*/
func reversePrint2(head *ListNode) []int {
    // base case
    if head == nil {
        return []int{}
    }

    // 把子链表翻转的结果算出来，示例 [3,2]
    subProblem := reversePrint2(head.Next)
    // 把 head 的值接到子链表的翻转结果的尾部，示例 [3,2,1]
    return append(subProblem, head.Val)
}
```

</div></div>

<div data-tab-item="javascript" class="tab-item " data-tab-group="default"><div class="highlight">

```javascript
// 注意：javascript 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
// 本代码已经通过力扣的测试用例，应该可直接成功提交。

/**
 * 用「遍历」的思路写递归解法
 */

var reversePrint = function(head) {
  let len = 0; // 记录链表长度
  let res = []; // 结果数组
  let p = 0; // 结果数组中的指针

  // 递归遍历单链表
  const traverse = (head) => {
    if (!head) {
      // 到达链表尾部，此时知道了链表的总长度
      // 创建结果数组
      res = new Array(len);
      return;
    }
    len++;
    traverse(head.next);
    // 后序位置，可以倒序操作链表
    res[p] = head.val;
    p++;
  };

  traverse(head);
  return res;
};


/**
 * 用「分解问题」的思路写递归解法
 * 因为 Java 的 int[] 数组不支持 add 相关的操作，所以我们把返回值修改成 List
 * 定义：输入一个单链表，返回该链表翻转的值，示例 1->2->3
 */
var reversePrint2 = function(head) {
  // base case
  if (!head) {
    return new LinkedList();
  }

  // 把子链表翻转的结果算出来，示例 [3,2]
  let subProblem = reversePrint2(head.next);
  // 把 head 的值接到子链表的翻转结果的尾部，示例 [3,2,1]
  subProblem.push(head.val);
  return subProblem;
};
```

</div></div>
</div></div>

<hr /><details open hint-container details><summary style="font-size: medium"><strong>🌈🌈 算法可视化 🌈🌈</strong></summary><div id="data_cong-wei-dao-tou-da-yin-lian-biao-lcof" data="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"></div><div class="resizable aspect-ratio-container" style="height: 100%;">
<div id="iframe_cong-wei-dao-tou-da-yin-lian-biao-lcof"></div></div>
</details><hr /><br />

</div>

</details>
</div>

