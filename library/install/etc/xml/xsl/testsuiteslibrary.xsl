<?xml version="1.0" encoding="UTF-8"?>
<xsl:stylesheet version="1.0" xmlns:xsl="http://www.w3.org/1999/XSL/Transform">
<xsl:template match="/">
<html>
<head>
  <link rel="stylesheet" type="text/css" href="css/testsuiteslibrary.css" />
</head>
<body>
<xsl:variable name="testsuitelibrary_name"><xsl:value-of select="/TestSuitesLibrary/@name" /></xsl:variable>
<div id="Container" class="Container">
  <div id="TestSuiteLibrary" class="TestSuiteLibrary">
    <div id="TestSuiteLibraryTitle" class="TestSuiteLibraryTitle">Custom <xsl:value-of select="/TestSuitesLibrary/@title" /> Testcases Library</div>
    <div>Index</div>
    <div class="Index">
      <xsl:for-each select="/TestSuitesLibrary/TestSuites/TestSuite">
      <ul>
        <xsl:variable name="testsuite_name"><xsl:value-of select="@name" /></xsl:variable>
        <li><a href="#{$testsuite_name}"> <xsl:value-of select="@title" /></a>
          <ul>
            <xsl:for-each select="TestCases/TestCase">
              <xsl:variable name="testcase_name"><xsl:value-of select="@name" /></xsl:variable>
              <li><a href="#{$testsuite_name}_{$testcase_name}"><xsl:value-of select="@title" /></a></li>
            </xsl:for-each>
          </ul>
        </li>
      </ul>
      </xsl:for-each>
    </div> <!-- Index -->
    <div class="Spacer">&#160;</div>
    <xsl:for-each select="/TestSuitesLibrary/TestSuites/TestSuite">
    <div class="TestSuite">
      <div class="Spacer">&#160;</div>
      <xsl:variable name="testsuite_title"><xsl:value-of select="@name" /></xsl:variable>
      <div class="TestSuiteTitle"><a name="{$testsuite_title}" class="TestSuiteTitle"><xsl:value-of select="@title" /></a></div>
      <div class="TestSuiteDescription"><xsl:value-of select="Description" /></div>
      <xsl:for-each select="TestCases/TestCase">
      <div class="TestCase">
        <div class="TestCaseRow">



          <xsl:variable name="testcase_title"><xsl:value-of select="@name" /></xsl:variable>
          <div class="TestCaseHeadFirst">
              <a name="#{$testsuite_title}_{$testcase_title}" class="TestCaseHeadFirst" ><xsl:value-of select="@title" />&#160;(Responsible:&#160;<xsl:value-of select="Rationale/@responsible" />)</a>
          </div> <!--TestCaseHeadFirst-->




        </div> <!--TestCaseRow-->





        <div class="TestCaseRow">

                    <div class="TestCaseHeadType">ID:</div>
                    <div class="TestCaseHeadValue"><xsl:value-of select="@id" /></div>

        </div> <!--TestCaseRow-->

        <div class="TestCaseRow">
          <div class="TestCaseHeadType">Benchmark:</div>
          <xsl:variable name="benchmark_name"><xsl:value-of select="@benchmark" /></xsl:variable>


          <div class="TestCaseNestedValue"><a href="benchmarklibrary.xml#{$benchmark_name}" class="TestCaseCell"><xsl:value-of select="@benchmark" /></a></div>
        </div> <!--TestCaseRow-->








        <div class="TestCaseRow">



          <div class="TestCaseHeadType">Applies To:</div>
          <div class="TestCaseHeadValue">
            <xsl:choose>
              <xsl:when test="not (Inject/Classes/Class[1])">
                Any Host
              </xsl:when>
              <xsl:otherwise>
                Any target host tagged with class
                <xsl:for-each select="Inject/Classes/Class">
                  &#160;&quot;<xsl:value-of select="." />&quot;
                </xsl:for-each> 
              </xsl:otherwise>
            </xsl:choose>
          </div> <!-- TestCaseCell -->


        </div> <!--TestCaseRow-->
        <div class="TestCaseRow">

          <div class="TestCaseNested">
                  <div class="TestCaseNestedRow">


          <div class="TestCaseNestedHeadType">Dependencies:</div>
          <div class="TestCaseNestedHeadValue">
            <xsl:choose>
             <xsl:when test="not (Dependencies/Dependency[1])">
                None
              </xsl:when>
              <xsl:otherwise>
                <div class="Dependencies">
                  <div class="DependenciesRow">
                    <div class="DependenciesHeadType">Type</div>
                    <div class="DependenciesHeadValue">Dependency</div>
                  </div> <!-- DependenciesRow -->
                  <xsl:for-each select="Dependencies/Dependency">
                  <div class="DependenciesRow">
                    <div class="DependenciesCellType"><xsl:value-of select="@type" /></div>
                    <xsl:variable name="dependency_testcase"><xsl:value-of select="@name" /></xsl:variable>
                    <xsl:variable name="dependency_testsuite"><xsl:value-of select="@testsuite" /></xsl:variable>
                    <div class="DependenciesCellType">
                      <a href="#{$dependency_testsuite}_{$dependency_testcase}"><xsl:value-of select="@title" /></a>
                    </div> <!-- DependenciesCellType -->
                  </div> <!-- DependenciesRow -->
                  </xsl:for-each>
                </div> <!-- Dependencies -->
              </xsl:otherwise>
            </xsl:choose>
          </div> <!-- TestCaseCell -->

</div> <!-- TestCaseNestedRow -->
</div> <!-- TestCaseNested -->



        </div> <!--TestCaseRow-->




        <div class="TestCaseRow">




          <div class="TestCaseCellFull"><xsl:value-of select="Description" /></div>






        </div> <!--TestCaseRow-->


        <div class="TestCaseRow">
          <div class="TestCaseCellFull">



                <div class="Tests">
                  <div class="TestsRow">
                    <div class="TestsHeadType">Test</div>
                    <div class="TestsHeadValue">TestSteps</div>
                  </div> <!-- TestsRow -->
                  <xsl:for-each select="Tests/Test">
                  <div class="TestsRow">
                    <div class="TestsCellType"><xsl:value-of select="@title" /></div>
                    <div class="TestsCellValue">


                  <table class="TestSteps">
                    <tr>
                      <xsl:for-each select="TestSteps/TestStep[1]/@*">
                        <th><xsl:value-of select="name()" /></th>
                      </xsl:for-each>
                    </tr>
                    <xsl:for-each select="TestSteps/TestStep">
                    <tr>
                      <xsl:for-each select="@*">
                        <td><xsl:value-of select="." /></td>
                      </xsl:for-each>
                    </tr>
                    </xsl:for-each>
                  </table>


                    </div> <!-- TestsCellValue -->

<div class="TestsCellFull">



                  <table class="TestParameters">
                    <tr>
                      <xsl:for-each select="Parameters/Parameter[1]/@*">
                        <th><xsl:value-of select="name()" /></th>
                      </xsl:for-each>
                    </tr>
                    <xsl:for-each select="Parameters/Parameter">
                    <tr>
                      <xsl:for-each select="@*">
                        <td><xsl:value-of select="." /></td>
                      </xsl:for-each>
                    </tr>
                    </xsl:for-each>
                  </table>




</div>


                  </div> <!-- TestRow -->
                  </xsl:for-each>
                </div> <!-- Dependencies -->


          </div> <!-- TestCaseCell -->



        </div> <!--TestCaseRow-->











        <div class="TestCaseRow">



          <div class="TestCaseNested">
                  <div class="TestCaseNestedRow">

          <div class="TestCaseNestedHeadType">Rationale:</div>


          <div class="TestCaseNestedHeadValue">
            <table class="Rationale">
              <tr>
                <td class="RationaleHead">Health Threeshold:</td>
                <td class="RationaleFirstCol">
                  <xsl:choose>
                    <xsl:when test="not (Rationale/@minimum_health)">
                      0%
                    </xsl:when>
                    <xsl:otherwise>
                      <xsl:value-of select="Rationale/@minimum_health" />
                    </xsl:otherwise>
                  </xsl:choose>
                </td>
                <td class="RationaleSecondCol">
                  Health above this percentage means that the item should be considered unhealthy. Health below this percentage instead means that the item shuold be considered failed.
                </td>
              </tr>
              <tr>
                <td class="RationaleHead">Severity:</td>
                <td class="RationaleFirstCol">
                <xsl:choose>
                  <xsl:when test="contains (Rationale/@severity, 'Fatal')">
                    <td class="red"><xsl:value-of select="Rationale/@severity" /></td>
                  </xsl:when>
                  <xsl:when test="contains (Rationale/@severity, 'Critical')">
                    <td class="amber"><xsl:value-of select="Rationale/@severity" /></td>
                  </xsl:when>
                  <xsl:otherwise>
                    <td class="green"><xsl:value-of select="Rationale/@severity" /></td>
                  </xsl:otherwise>
                </xsl:choose>
                </td>
                <td class="RationaleSecondCol">
                  Fatal: a failure of this item can compromise the whole tested environment<br/>Critical: a failure of this item can be considered a problem to be fixed but does not block the whole tested environment.
                </td>
              </tr>
            </table>
          </div> <!--TestCaseCell-->

</div> <!-- TestCaseNestedRow -->
</div> <!-- TestCaseNested -->



        </div> <!--TestCaseRow-->

      </div> <!--TestCase-->
      </xsl:for-each>
    </div> <!-- TestSuite -->
    </xsl:for-each>
  </div> <!-- TestSuiteLibrary -->
</div> <!-- Container -->
</body>
</html>
</xsl:template>
</xsl:stylesheet>
